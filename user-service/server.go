package main

import (
	"log"
	"net/http"
	"user-service/graph"
	"user-service/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/graphql", srv)
	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

	log.Println("connect to http://localhost:8081/playground for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
