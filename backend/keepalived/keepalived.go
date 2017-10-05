// Package keepalived provides client keepalive ("aliveness") monitoring
// functionality for Sensu. The principle mechanism responsible for
// keepalive monitoring is Keepalived.
//
// When a sensu-agent process connects, it begins sending keepalive
// (transport.MessageTypeKeepalive) messages across the transport. These
// are published to the keepalive message topic (messaging.TopicKeepalive)
// to which a Keepalived subscribes.
//
// For more information on how Keepalives are handled, see the
// documentation for Keepalived.
package keepalived

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/sensu/sensu-go/backend/messaging"
	"github.com/sensu/sensu-go/backend/store"
	"github.com/sensu/sensu-go/types"
)

const (
	// DefaultHandlerCount is the default number of goroutines dedicated to
	// handling keepalive events.
	DefaultHandlerCount = 10

	// DefaultKeepaliveTimeout is the amount of time we consider a Keepalive
	// valid for.
	DefaultKeepaliveTimeout = 120 // seconds
)

// MonitorFactoryFunc takes an entity and returns a Monitor. Keepalived can
// take a MonitorFactoryFunc that stubs/mocks a Deregisterer and/or an
// EventCreator to make it easier to test.
type MonitorFactoryFunc func(e *types.Entity) *KeepaliveMonitor

// Keepalived is responsible for monitoring keepalive events and recording
// keepalives for entities.
//
// Each instance of Keepalived has a collection of KeepaliveMonitors that
// are individually responsible for monitoring the aliveness of a Sensu
// agent. They do this with limited coordination via the Store which
// somewhat complicates their behavior.
//
// When a Keepalive message is received via the MessageBus, Keepalived
// checks to see if there is an existing monitor for the Entity associated
// with that Keepalive. If none exists, it creates a new monitor and
// starts it. If a monitor exists, but has been stopped (see the documentation
// for KeepaliveMonitor), it will create and start a new monitor.
//
// Keepalived also periodically sweeps its monitor map for stopped monitors
// (representing clients that have been deregistered or have connected to
// another backend and not reconnected here) so that they can be GC'd.
type Keepalived struct {
	MessageBus            messaging.MessageBus
	HandlerCount          int
	Store                 store.Store
	DeregistrationHandler string
	MonitorFactory        MonitorFactoryFunc

	mu            *sync.Mutex
	monitors      map[string]*KeepaliveMonitor
	wg            *sync.WaitGroup
	keepaliveChan chan interface{}
	errChan       chan error
}

// Start starts the daemon, returning an error if preconditions for startup
// fail.
func (k *Keepalived) Start() error {
	if k.MessageBus == nil {
		return errors.New("no message bus found")
	}

	if k.Store == nil {
		return errors.New("no keepalive store found")
	}

	if k.MonitorFactory == nil {
		k.MonitorFactory = func(e *types.Entity) *KeepaliveMonitor {
			return &KeepaliveMonitor{
				Entity: e,
				Deregisterer: &Deregistration{
					Store:      k.Store,
					MessageBus: k.MessageBus,
				},
				EventCreator: &MessageBusEventCreator{
					MessageBus: k.MessageBus,
				},
				Store: k.Store,
			}
		}
	}

	k.keepaliveChan = make(chan interface{}, 10)
	err := k.MessageBus.Subscribe(messaging.TopicKeepalive, "keepalived", k.keepaliveChan)
	if err != nil {
		return err
	}

	if k.HandlerCount == 0 {
		k.HandlerCount = DefaultHandlerCount
	}

	k.mu = &sync.Mutex{}
	k.monitors = map[string]*KeepaliveMonitor{}

	if err := k.initFromStore(); err != nil {
		return err
	}

	k.startWorkers()

	k.startMonitorSweeper()

	k.errChan = make(chan error, 1)
	return nil
}

// Stop stops the daemon, returning an error if one was encountered during
// shutdown.
func (k *Keepalived) Stop() error {
	close(k.keepaliveChan)
	k.wg.Wait()
	for _, monitor := range k.monitors {
		go monitor.Stop()
	}
	k.MessageBus.Unsubscribe(messaging.TopicKeepalive, "keepalived")
	close(k.errChan)
	return nil
}

// Status returns nil if the Daemon is healthy, otherwise it returns an error.
func (k *Keepalived) Status() error {
	return nil
}

// Err returns a channel that the caller can use to listen for terminal errors
// indicating a premature shutdown of the Daemon.
func (k *Keepalived) Err() <-chan error {
	return k.errChan
}

func (k *Keepalived) initFromStore() error {
	// For which clients were we previously alerting?
	keepalives, err := k.Store.GetFailingKeepalives(context.TODO())
	if err != nil {
		return err
	}

	for _, keepalive := range keepalives {
		entityCtx := context.WithValue(context.TODO(), types.OrganizationKey, keepalive.Organization)
		entityCtx = context.WithValue(entityCtx, types.EnvironmentKey, keepalive.Environment)
		event, err := k.Store.GetEventByEntityCheck(entityCtx, keepalive.EntityID, "keepalive")
		if err != nil {
			return err
		}

		// if there's no event, the entity was deregistered/deleted.
		if event == nil {
			continue
		}

		// if another backend picked it up, it will be passing.
		if event.Check.Status == 0 {
			continue
		}

		// Recreate the monitor and reset its timer to alert when it's going to
		// timeout.
		monitor := k.MonitorFactory(event.Entity)
		monitor.Reset(keepalive.Time)
		k.monitors[keepalive.EntityID] = monitor
	}

	return nil
}

func (k *Keepalived) startWorkers() {
	k.wg = &sync.WaitGroup{}
	k.wg.Add(k.HandlerCount)

	for i := 0; i < k.HandlerCount; i++ {
		go k.processKeepalives()
	}
}

func (k *Keepalived) processKeepalives() {
	defer k.wg.Done()

	var (
		monitor *KeepaliveMonitor
		event   *types.Event
		ok      bool
	)

	for msg := range k.keepaliveChan {
		event, ok = msg.(*types.Event)
		if !ok {
			logger.Error("keepalived received non-Event on keepalive channel")
			continue
		}

		entity := event.Entity
		if err := entity.Validate(); err != nil {
			logger.WithError(err).Error("invalid keepalive event")
			continue
		}

		k.mu.Lock()
		monitor, ok = k.monitors[entity.ID]
		// create if it doesn't exist
		if !ok || monitor.IsStopped() {
			monitor = k.MonitorFactory(entity)
			monitor.Start()
			k.monitors[entity.ID] = monitor
		}
		k.mu.Unlock()

		if err := monitor.Update(event); err != nil {
			logger.WithError(err).Error("error monitoring entity")
		}
	}
}

// startMonitorSweeper spins off into oblivion if Keepalived is stopped until
// the monitors map is empty, and then the goroutine stops.
func (k *Keepalived) startMonitorSweeper() {
	go func() {
		timer := time.NewTimer(10 * time.Minute)
		for {
			<-timer.C
			for key, monitor := range k.monitors {
				if monitor.IsStopped() {
					k.mu.Lock()
					delete(k.monitors, key)
					k.mu.Unlock()
				}
			}
		}
	}()
}
