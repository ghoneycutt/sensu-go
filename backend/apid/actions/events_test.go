package actions

import (
	"context"
	"errors"
	"testing"

	"github.com/sensu/sensu-go/testing/mockstore"
	"github.com/sensu/sensu-go/testing/testutil"
	"github.com/sensu/sensu-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewEventController(t *testing.T) {
	assert := assert.New(t)

	store := &mockstore.MockStore{}
	eventController := NewEventController(store)

	assert.NotNil(eventController)
	assert.Equal(store, eventController.Store)
	assert.NotNil(eventController.Policy)
}

func TestEventQuery(t *testing.T) {
	defaultCtx := testutil.NewContext(testutil.ContextWithRules(
		types.FixtureRuleWithPerms(types.RuleTypeEvent, types.RulePermRead),
	))

	testCases := []struct {
		name        string
		ctx         context.Context
		events      []*types.Event
		params      QueryParams
		expectedLen int
		storeErr    error
		expectedErr error
	}{
		{
			name:        "No Params No Events",
			ctx:         defaultCtx,
			events:      []*types.Event{},
			params:      QueryParams{},
			expectedLen: 0,
			storeErr:    nil,
			expectedErr: nil,
		},
		{
			name: "No Params With Events",
			ctx:  defaultCtx,
			events: []*types.Event{
				types.FixtureEvent("entity1", "check1"),
				types.FixtureEvent("entity2", "check2"),
			},
			params:      QueryParams{},
			expectedLen: 2,
			storeErr:    nil,
			expectedErr: nil,
		},
		{
			name: "No Params With Only Create Access",
			ctx: testutil.NewContext(testutil.ContextWithRules(
				types.FixtureRuleWithPerms(types.RuleTypeEvent, types.RulePermCreate),
			)),
			events: []*types.Event{
				types.FixtureEvent("entity1", "check1"),
				types.FixtureEvent("entity2", "check2"),
			},
			params:      QueryParams{},
			expectedLen: 0,
			storeErr:    nil,
			expectedErr: nil,
		},
		{
			name: "Entity Param",
			ctx:  defaultCtx,
			events: []*types.Event{
				types.FixtureEvent("entity1", "check1"),
			},
			params: QueryParams{
				"entity": "entity1",
			},
			expectedLen: 1,
			storeErr:    nil,
			expectedErr: nil,
		},
		{
			name:   "Store Failure",
			ctx:    defaultCtx,
			events: nil,
			params: QueryParams{
				"entity": "entity1",
			},
			expectedLen: 0,
			storeErr:    errors.New(""),
			expectedErr: NewError(InternalErr, errors.New("")),
		},
	}

	for _, tc := range testCases {
		store := &mockstore.MockStore{}
		eventController := NewEventController(store)

		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			// Mock store methods
			store.On("GetEvents", tc.ctx).Return(tc.events, tc.storeErr)
			store.On("GetEventsByEntity", tc.ctx, mock.Anything).Return(tc.events, tc.storeErr)

			// Exec Query
			results, err := eventController.Query(tc.ctx, tc.params)

			// Assert
			assert.EqualValues(tc.expectedErr, err)
			assert.Len(results, tc.expectedLen)
		})
	}
}

func TestEventFind(t *testing.T) {
	defaultCtx := testutil.NewContext(testutil.ContextWithRules(
		types.FixtureRuleWithPerms(types.RuleTypeEvent, types.RulePermRead),
	))

	testCases := []struct {
		name            string
		ctx             context.Context
		event           *types.Event
		params          QueryParams
		expected        bool
		expectedErrCode ErrCode
	}{
		{
			name:            "No Params",
			ctx:             defaultCtx,
			params:          QueryParams{},
			expected:        false,
			expectedErrCode: InvalidArgument,
		},
		{
			name: "Only Entity Param",
			ctx:  defaultCtx,
			params: QueryParams{
				"entity": "entity1",
			},
			expected:        false,
			expectedErrCode: InvalidArgument,
		},
		{
			name: "Only Check Param",
			ctx:  defaultCtx,
			params: QueryParams{
				"check": "check1",
			},
			expected:        false,
			expectedErrCode: InvalidArgument,
		},
		{
			name:  "Found",
			ctx:   defaultCtx,
			event: types.FixtureEvent("entity1", "check1"),
			params: QueryParams{
				"entity": "entity1",
				"check":  "check1",
			},
			expected:        true,
			expectedErrCode: 0,
		},
		{
			name:  "Not Found",
			ctx:   defaultCtx,
			event: nil,
			params: QueryParams{
				"entity": "entity1",
				"check":  "check1",
			},
			expected:        false,
			expectedErrCode: NotFound,
		},
		{
			name: "No Read Permission",
			ctx: testutil.NewContext(testutil.ContextWithRules(
				types.FixtureRuleWithPerms(types.RuleTypeEvent, types.RulePermCreate),
			)),
			event: types.FixtureEvent("entity1", "check1"),
			params: QueryParams{
				"entity": "entity1",
				"check":  "check1",
			},
			expected:        false,
			expectedErrCode: NotFound,
		},
	}

	for _, tc := range testCases {
		store := &mockstore.MockStore{}
		eventController := NewEventController(store)

		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			// Mock store methods
			store.
				On("GetEventByEntityCheck", tc.ctx, mock.Anything, mock.Anything).
				Return(tc.event, nil)

			// Exec Query
			result, err := eventController.Find(tc.ctx, tc.params)

			inferErr, ok := err.(Error)
			if ok {
				assert.Equal(tc.expectedErrCode, inferErr.Code)
			} else {
				assert.NoError(err)
			}
			assert.Equal(tc.expected, result != nil, "expects Find() to return an event")
		})
	}
}

func TestEventDestroy(t *testing.T) {
	defaultCtx := testutil.NewContext(testutil.ContextWithRules(
		types.FixtureRuleWithPerms(types.RuleTypeEvent, types.RulePermDelete),
	))

	testCases := []struct {
		name            string
		ctx             context.Context
		event           *types.Event
		params          QueryParams
		expected        bool
		expectedErrCode ErrCode
	}{
		{
			name:            "No Params",
			ctx:             defaultCtx,
			params:          QueryParams{},
			expectedErrCode: InvalidArgument,
		},
		{
			name: "Only Entity Param",
			ctx:  defaultCtx,
			params: QueryParams{
				"entity": "entity1",
			},
			expectedErrCode: InvalidArgument,
		},
		{
			name: "Only Check Param",
			ctx:  defaultCtx,
			params: QueryParams{
				"check": "check1",
			},
			expectedErrCode: InvalidArgument,
		},
		{
			name:  "Delete",
			ctx:   defaultCtx,
			event: types.FixtureEvent("entity1", "check1"),
			params: QueryParams{
				"entity": "entity1",
				"check":  "check1",
			},
			expectedErrCode: 0,
		},
		{
			name:  "Not Found",
			ctx:   defaultCtx,
			event: nil,
			params: QueryParams{
				"entity": "entity1",
				"check":  "check1",
			},
			expectedErrCode: NotFound,
		},
		{
			name: "No Delete Permission",
			ctx: testutil.NewContext(testutil.ContextWithRules(
				types.FixtureRuleWithPerms(types.RuleTypeEvent, types.RulePermCreate),
			)),
			event: types.FixtureEvent("entity1", "check1"),
			params: QueryParams{
				"entity": "entity1",
				"check":  "check1",
			},
			expectedErrCode: NotFound,
		},
	}

	for _, tc := range testCases {
		store := &mockstore.MockStore{}
		eventController := NewEventController(store)

		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			// Mock store methods
			store.
				On("GetEventByEntityCheck", tc.ctx, mock.Anything, mock.Anything).
				Return(tc.event, nil)
			store.
				On("DeleteEventByEntityCheck", tc.ctx, mock.Anything, mock.Anything).
				Return(nil)

			// Exec Query
			err := eventController.Destroy(tc.ctx, tc.params)

			inferErr, ok := err.(Error)
			if ok {
				assert.Equal(tc.expectedErrCode, inferErr.Code)
			} else {
				assert.NoError(err)
			}
		})
	}
}

func TestEventUpdate(t *testing.T) {
	defaultCtx := testutil.NewContext(
		testutil.ContextWithRules(
			types.FixtureRuleWithPerms(types.RuleTypeEvent, types.RulePermUpdate),
		),
	)
	wrongPermsCtx := testutil.NewContext(
		testutil.ContextWithRules(
			types.FixtureRuleWithPerms(types.RuleTypeEvent, types.RulePermRead),
		),
	)

	badEvent := types.FixtureEvent("entity1", "check1")
	badEvent.Check.Config.Name = "!@#!#$@#^$%&$%&$&$%&%^*%&(%@###"

	testCases := []struct {
		name            string
		ctx             context.Context
		argument        *types.Event
		fetchResult     *types.Event
		fetchErr        error
		updateErr       error
		expectedErr     bool
		expectedErrCode ErrCode
	}{
		{
			name:        "Updated",
			ctx:         defaultCtx,
			argument:    types.FixtureEvent("entity1", "check1"),
			fetchResult: types.FixtureEvent("entity1", "check1"),
			expectedErr: false,
		},
		{
			name:            "Does Not Exist",
			ctx:             defaultCtx,
			argument:        types.FixtureEvent("entity1", "check1"),
			fetchResult:     nil,
			expectedErr:     true,
			expectedErrCode: NotFound,
		},
		{
			name:            "Store Err on Update",
			ctx:             defaultCtx,
			argument:        types.FixtureEvent("entity1", "check1"),
			fetchResult:     types.FixtureEvent("entity1", "check1"),
			updateErr:       errors.New("dunno"),
			expectedErr:     true,
			expectedErrCode: InternalErr,
		},
		{
			name:            "Store Err on Fetch",
			ctx:             defaultCtx,
			argument:        types.FixtureEvent("entity1", "check1"),
			fetchResult:     types.FixtureEvent("entity1", "check1"),
			fetchErr:        errors.New("dunno"),
			expectedErr:     true,
			expectedErrCode: InternalErr,
		},
		{
			name:            "No Permission",
			ctx:             wrongPermsCtx,
			argument:        types.FixtureEvent("entity1", "check1"),
			fetchResult:     types.FixtureEvent("entity1", "check1"),
			expectedErr:     true,
			expectedErrCode: PermissionDenied,
		},
		{
			name:            "Validation Error",
			ctx:             defaultCtx,
			argument:        badEvent,
			fetchResult:     types.FixtureEvent("entity1", "check1"),
			expectedErr:     true,
			expectedErrCode: InvalidArgument,
		},
	}

	for _, tc := range testCases {
		store := &mockstore.MockStore{}
		actions := NewEventController(store)

		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			// Mock store methods
			store.
				On("GetEventByEntityCheck", mock.Anything, mock.Anything, mock.Anything).
				Return(tc.fetchResult, tc.fetchErr)
			store.
				On("UpdateEvent", mock.Anything, mock.Anything).Return(tc.updateErr)

			// Exec Query
			err := actions.Update(tc.ctx, *tc.argument)

			if tc.expectedErr {
				inferErr, ok := err.(Error)
				if ok {
					assert.Equal(tc.expectedErrCode, inferErr.Code)
				} else {
					assert.Error(err)
					assert.FailNow("Given was not of type 'Error'")
				}
			} else {
				assert.NoError(err)
			}
		})
	}
}

func TestEventCreate(t *testing.T) {
	defaultCtx := testutil.NewContext(
		testutil.ContextWithRules(
			types.FixtureRuleWithPerms(types.RuleTypeEvent, types.RulePermCreate),
		),
	)
	wrongPermsCtx := testutil.NewContext(
		testutil.ContextWithRules(
			types.FixtureRuleWithPerms(types.RuleTypeEvent, types.RulePermRead),
		),
	)

	badEvent := types.FixtureEvent("entity1", "check1")
	badEvent.Check.Config.Name = "!@#!#$@#^$%&$%&$&$%&%^*%&(%@###"

	testCases := []struct {
		name            string
		ctx             context.Context
		argument        *types.Event
		fetchResult     *types.Event
		fetchErr        error
		createErr       error
		expectedErr     bool
		expectedErrCode ErrCode
	}{
		{
			name:        "Created",
			ctx:         defaultCtx,
			argument:    types.FixtureEvent("entity1", "check1"),
			expectedErr: false,
		},
		{
			name:        "Already Exists",
			ctx:         defaultCtx,
			argument:    types.FixtureEvent("entity1", "check1"),
			fetchResult: types.FixtureEvent("entity1", "check1"),
			expectedErr: false,
		},
		{
			name:            "Store Err on Create",
			ctx:             defaultCtx,
			argument:        types.FixtureEvent("entity1", "check1"),
			createErr:       errors.New("dunno"),
			expectedErr:     true,
			expectedErrCode: InternalErr,
		},
		{
			name:            "Store Err on Fetch",
			ctx:             defaultCtx,
			argument:        types.FixtureEvent("entity1", "check1"),
			fetchErr:        errors.New("dunno"),
			expectedErr:     true,
			expectedErrCode: InternalErr,
		},
		{
			name:            "No Permission",
			ctx:             wrongPermsCtx,
			argument:        types.FixtureEvent("entity1", "check1"),
			expectedErr:     true,
			expectedErrCode: PermissionDenied,
		},
		{
			name:            "Validation Error",
			ctx:             defaultCtx,
			argument:        badEvent,
			expectedErr:     true,
			expectedErrCode: InvalidArgument,
		},
	}

	for _, tc := range testCases {
		store := &mockstore.MockStore{}
		actions := NewEventController(store)

		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			// Mock store methods
			store.
				On("GetEventByEntityCheck", mock.Anything, mock.Anything, mock.Anything).
				Return(tc.fetchResult, tc.fetchErr)
			store.
				On("UpdateEvent", mock.Anything, mock.Anything).Return(tc.createErr)

			// Exec Query
			err := actions.Create(tc.ctx, *tc.argument)
			if tc.expectedErr {
				inferErr, ok := err.(Error)
				if ok {
					assert.Equal(tc.expectedErrCode, inferErr.Code)
				} else {
					assert.Error(err)
					assert.FailNow("Given was not of type 'Error'")
				}
			} else {
				assert.NoError(err)
			}
		})
	}
}
