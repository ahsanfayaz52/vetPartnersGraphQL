package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
	"vetPartnersGraphQL/config"
	"vetPartnersGraphQL/graph"
	"vetPartnersGraphQL/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://%s:%s/ for GraphQL playground", viper.GetString(config.ServerHost), viper.GetString(config.ServerPort))
	log.Fatal(http.ListenAndServe(":"+viper.GetString(config.ServerPort), nil))
}
