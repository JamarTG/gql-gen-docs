package main

import (
	"example/graph"
	"example/graph/model"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"


func main() {

	r := &graph.Resolver{
		TodosData: []*model.Todo{
			{
				ID:   "1",
				Text: "Learn GraphQL",
				Done: false,
				User: &model.User{
					ID:   "u1",
					Name: "Jamari",
				},
			},
			{
				ID:   "2",
				Text: "Build API",
				Done: false,
				User: &model.User{
					ID:   "u1",
					Name: "Jamari",
				},
			},
		},
		Users: []*model.User{
			{
				ID:   "u1",
				Name: "Jamari",
			},
		},
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: r}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
