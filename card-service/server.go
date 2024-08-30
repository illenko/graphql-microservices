// card_service/server.go
package main

import (
	"card-service/graph"
	"card-service/graph/generated"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/graphql", srv)
	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

	log.Println("connect to http://localhost:8082/playground for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
