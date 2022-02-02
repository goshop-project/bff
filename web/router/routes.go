package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"goshop.dev/bff/graph"
	"goshop.dev/bff/graph/generated"
	"goshop.dev/bff/tools/valid"
)

type RoutesConfig struct {
	GraphQL           string `default:"/q/query" valid:"pathdis" yaml:"graphql,omitempty"`
	GraphQLPlayground string `default:"/q"       valid:"pathdis" yaml:"graphqlplayground,omitempty"`
}

func (rc *RoutesConfig) Route(router Router) error {
	var errs valid.Errors

	if rc.GraphQL == "-" {
		// skip graphql
	} else if !valid.IsPath(rc.GraphQL) {
		// invalid graphql path
		errs.AppendFieldError("graphql", "path", nil)
	} else {

		if rc.GraphQLPlayground == "-" {
			// skip playground
		} else if !valid.IsPath(rc.GraphQLPlayground) {
			// invalid playground path
			errs.AppendFieldError("graphqlplayground", "path", nil)
		} else {
			// handle playground
			h := playground.Handler("GraphQL playground", rc.GraphQL)
			router.Handle(rc.GraphQLPlayground, h)
		}

		// handle graphql
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
		router.Handle(rc.GraphQL, srv)
	}

	return errs.AsError()
}
