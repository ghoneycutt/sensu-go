// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	graphql1 "github.com/graphql-go/graphql"
	graphql "github.com/sensu/sensu-go/graphql"
)

// Schema supplies the root types of each type of operation, query,
// mutation (optional), and subscription (optional).
var Schema = graphql.NewType("Schema", graphql.SchemaKind)

// RegisterSchema registers schema description with given service.
func RegisterSchema(svc graphql.Service) {
	svc.RegisterSchema(_SchemaDesc)
}
func _SchemaConfigFn() graphql1.SchemaConfig {
	return graphql1.SchemaConfig{
		Mutation: graphql.Object("Mutation"),
		Query:    graphql.Object("Query"),
	}
}

// describe schema's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _SchemaDesc = graphql.SchemaDesc{Config: _SchemaConfigFn}

// QueryChecksFieldResolver implement to resolve requests for the Query's checks field.
type QueryChecksFieldResolver interface {
	// Checks implements response to request for checks field.
	Checks(p graphql.ResolveParams) (interface{}, error)
}

//
// QueryFieldResolvers represents a collection of methods whose products represent the
// response values of the 'Query' type.
//
// == Example SDL
//
//   """
//   Dog's are not hooman.
//   """
//   type Dog implements Pet {
//     "name of this fine beast."
//     name:  String!
//
//     "breed of this silly animal; probably shibe."
//     breed: [Breed]
//   }
//
// == Example generated interface
//
//   // DogResolver ...
//   type DogFieldResolvers interface {
//     DogNameFieldResolver
//     DogBreedFieldResolver
//
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
// == Example implementation ...
//
//   // DogResolver implements DogFieldResolvers interface
//   type DogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *DogResolver) Name(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *DogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
//
type QueryFieldResolvers interface {
	QueryChecksFieldResolver

	// IsTypeOf is used to determine if a given value is associated with the Query type
	IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
}

// QueryAliases implements all methods on QueryFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := v.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
//
type QueryAliases struct{}

// Checks implements response to request for 'checks' field.
func (_ QueryAliases) Checks(p graphql.ResolveParams) (interface{}, error) {
	return graphql.DefaultResolver(p.Source, p.Info.FieldName)
}

// QueryType The query root of Sensu's GraphQL interface.
var QueryType = graphql.NewType("Query", graphql.ObjectKind)

// RegisterQuery registers Query object type with given service.
func RegisterQuery(svc graphql.Service, impl QueryFieldResolvers) {
	svc.RegisterObject(_ObjTypeQueryDesc, impl)
}
func _ObjTypeQueryChecksHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(QueryChecksFieldResolver)
	return resolver.Checks
}

func _ObjTypeQueryConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "The query root of Sensu's GraphQL interface.",
		Fields: graphql1.Fields{"checks": &graphql1.Field{
			Args:              graphql1.FieldConfigArgument{},
			DeprecationReason: "",
			Description:       "self descriptive",
			Name:              "checks",
			Type:              graphql1.NewList(graphql.OutputType("Check")),
		}},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see QueryFieldResolvers.")
		},
		Name: "Query",
	}
}

// describe Query's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjTypeQueryDesc = graphql.ObjectDesc{
	Config:        _ObjTypeQueryConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{"Checks": _ObjTypeQueryChecksHandler},
}
