// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	graphql1 "github.com/graphql-go/graphql"
	graphql "github.com/sensu/sensu-go/graphql"
)

// AssetIDFieldResolver implement to resolve requests for the Asset's id field.
type AssetIDFieldResolver interface {
	// ID implements response to request for id field.
	ID(p graphql.ResolveParams) (interface{}, error)
}

// AssetNamespaceFieldResolver implement to resolve requests for the Asset's namespace field.
type AssetNamespaceFieldResolver interface {
	// Namespace implements response to request for namespace field.
	Namespace(p graphql.ResolveParams) (interface{}, error)
}

// AssetNameFieldResolver implement to resolve requests for the Asset's name field.
type AssetNameFieldResolver interface {
	// Name implements response to request for name field.
	Name(p graphql.ResolveParams) (string, error)
}

// AssetUrlFieldResolver implement to resolve requests for the Asset's url field.
type AssetUrlFieldResolver interface {
	// Url implements response to request for url field.
	Url(p graphql.ResolveParams) (string, error)
}

// AssetSha512FieldResolver implement to resolve requests for the Asset's sha512 field.
type AssetSha512FieldResolver interface {
	// Sha512 implements response to request for sha512 field.
	Sha512(p graphql.ResolveParams) (string, error)
}

// AssetFiltersFieldResolver implement to resolve requests for the Asset's filters field.
type AssetFiltersFieldResolver interface {
	// Filters implements response to request for filters field.
	Filters(p graphql.ResolveParams) (string, error)
}

// AssetOrganizationFieldResolver implement to resolve requests for the Asset's organization field.
type AssetOrganizationFieldResolver interface {
	// Organization implements response to request for organization field.
	Organization(p graphql.ResolveParams) (string, error)
}

//
// AssetFieldResolvers represents a collection of methods whose products represent the
// response values of the 'Asset' type.
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
type AssetFieldResolvers interface {
	AssetIDFieldResolver
	AssetNamespaceFieldResolver
	AssetNameFieldResolver
	AssetUrlFieldResolver
	AssetSha512FieldResolver
	AssetFiltersFieldResolver
	AssetOrganizationFieldResolver

	// IsTypeOf is used to determine if a given value is associated with the Asset type
	IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
}

// AssetAliases implements all methods on AssetFieldResolvers interface by using reflection to
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
type AssetAliases struct{}

// ID implements response to request for 'id' field.
func (_ AssetAliases) ID(p graphql.ResolveParams) (interface{}, error) {
	return graphql.DefaultResolver(p.Source, p.Info.FieldName)
}

// Namespace implements response to request for 'namespace' field.
func (_ AssetAliases) Namespace(p graphql.ResolveParams) (interface{}, error) {
	return graphql.DefaultResolver(p.Source, p.Info.FieldName)
}

// Name implements response to request for 'name' field.
func (_ AssetAliases) Name(p graphql.ResolveParams) (string, error) {
	return graphql.DefaultResolver(p.Source, p.Info.FieldName)
}

// Url implements response to request for 'url' field.
func (_ AssetAliases) Url(p graphql.ResolveParams) (string, error) {
	return graphql.DefaultResolver(p.Source, p.Info.FieldName)
}

// Sha512 implements response to request for 'sha512' field.
func (_ AssetAliases) Sha512(p graphql.ResolveParams) (string, error) {
	return graphql.DefaultResolver(p.Source, p.Info.FieldName)
}

// Filters implements response to request for 'filters' field.
func (_ AssetAliases) Filters(p graphql.ResolveParams) (string, error) {
	return graphql.DefaultResolver(p.Source, p.Info.FieldName)
}

// Organization implements response to request for 'organization' field.
func (_ AssetAliases) Organization(p graphql.ResolveParams) (string, error) {
	return graphql.DefaultResolver(p.Source, p.Info.FieldName)
}

// AssetType Asset defines an asset agents install as a dependency for a check.
var AssetType = graphql.NewType("Asset", graphql.ObjectKind)

// RegisterAsset registers Asset object type with given service.
func RegisterAsset(svc graphql.Service, impl AssetFieldResolvers) {
	svc.RegisterObject(_ObjTypeAssetDesc, impl)
}
func _ObjTypeAssetIDHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(AssetIDFieldResolver)
	return resolver.ID
}

func _ObjTypeAssetNamespaceHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(AssetNamespaceFieldResolver)
	return resolver.Namespace
}

func _ObjTypeAssetNameHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(AssetNameFieldResolver)
	return resolver.Name
}

func _ObjTypeAssetUrlHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(AssetUrlFieldResolver)
	return resolver.Url
}

func _ObjTypeAssetSha512Handler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(AssetSha512FieldResolver)
	return resolver.Sha512
}

func _ObjTypeAssetFiltersHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(AssetFiltersFieldResolver)
	return resolver.Filters
}

func _ObjTypeAssetOrganizationHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(AssetOrganizationFieldResolver)
	return resolver.Organization
}

func _ObjTypeAssetConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "Asset defines an asset agents install as a dependency for a check.",
		Fields: graphql1.Fields{
			"filters": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Filters are a collection of sensu queries, used by the system to determine\nif the asset should be installed. If more than one filter is present the\nqueries are joined by the \"AND\" operator.",
				Name:              "filters",
				Type:              graphql1.String,
			},
			"id": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "id",
				Type:              graphql1.NewNonNull(graphql.OutputType("ID")),
			},
			"name": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Name is the unique identifier for an asset",
				Name:              "name",
				Type:              graphql1.String,
			},
			"namespace": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "namespace",
				Type:              graphql1.NewNonNull(graphql.OutputType("Namespace")),
			},
			"organization": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Organization indicates to which org an asset belongs to",
				Name:              "organization",
				Type:              graphql1.String,
			},
			"sha512": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Sha512 is the SHA-512 checksum of the asset",
				Name:              "sha512",
				Type:              graphql1.String,
			},
			"url": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "URL is the location of the asset",
				Name:              "url",
				Type:              graphql1.String,
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see AssetFieldResolvers.")
		},
		Name: "Asset",
	}
}

// describe Asset's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjTypeAssetDesc = graphql.ObjectDesc{
	Config: _ObjTypeAssetConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"Filters":      _ObjTypeAssetFiltersHandler,
		"ID":           _ObjTypeAssetIDHandler,
		"Name":         _ObjTypeAssetNameHandler,
		"Namespace":    _ObjTypeAssetNamespaceHandler,
		"Organization": _ObjTypeAssetOrganizationHandler,
		"Sha512":       _ObjTypeAssetSha512Handler,
		"Url":          _ObjTypeAssetUrlHandler,
	},
}
