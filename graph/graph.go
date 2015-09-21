package graph

import (
	"fmt"
	"github.com/tallstreet/graphql"
	"github.com/tallstreet/graphql/executor/resolver"
	"github.com/tallstreet/graphql/schema"
	"golang.org/x/net/context"
)

type Graph struct {
	nextId   int  
	Todos    map[int]*TodoNode
	Users    map[string]*User
}

func NewGraph() *Graph {
	graph := &Graph{
		0,
		make(map[int]*TodoNode),
		make(map[string]*User),
	}
	
	graph.Users["me"] = &User{
		"me",
		new(TodoConnection),
	}
	
	graph.AddToDo(graph.Users["me"], "Taste Javascript", false)
	graph.AddToDo(graph.Users["me"], "Buy a unicorn", false)

	return graph
}

func (graph *Graph) AddToDo(user *User, text string, complete bool) {
	
	todo := &TodoNode {
		fmt.Sprintf("%s", graph.nextId),
		text, 
		complete,
	};
	graph.Todos[graph.nextId] = todo
	graph.nextId += 1
	user.addToDo(todo)
	
}

func (graph *Graph) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "To Dos",
		Description: "A ToDo list App",
		Fields: schema.GraphQLFieldSpecMap{
			"viewer": {
				Name:        "viewer",
				Description: "A To Do user",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					g := graph.Users["me"]

					if g != nil {
						return r.Resolve(ctx, g, f)
					}
					return nil, fmt.Errorf("User not found")
				},
				IsRoot: true,
			},
		},
	}
}
