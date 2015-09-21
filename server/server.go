package server

import (
	"log"
	"net/http"

	"github.com/tallstreet/graphql/executor"
	"github.com/tallstreet/graphql/handler"
	"github.com/tallstreet/graphql/schema"
	"github.com/tallstreet/todographqlgo/graph"
)

type App struct {
	address string
}

var Application *App

func NewApp(address string) *App {
	app := new(App)
	app.address = address

	return app
}

func (app *App) RunServer() {
	g := graph.NewGraph()

	schema := schema.New()
	schema.Register(g.Todos[0])
	schema.Register(g.Users["me"])
	schema.Register(g)

	executor := executor.New(schema)
	handler := handler.New(executor)
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	log.Fatalln(http.ListenAndServe(app.address, mux))
}
