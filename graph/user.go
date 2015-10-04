package graph

import (
	"github.com/tallstreet/graphql"
	"github.com/tallstreet/graphql/executor/resolver"
	"github.com/tallstreet/graphql/schema"
	"golang.org/x/net/context"
)

type User struct {
	Id    string
	AnyTodos *TodoConnection
	CompletedTodos *TodoConnection
	ActiveTodos *TodoConnection
}

func (user *User) addToDo(todo *TodoEdge) {
	user.AnyTodos.addTodo(todo)
	if todo.Node.Completed {
		user.CompletedTodos.addTodo(todo)
	} else {
		user.ActiveTodos.addTodo(todo)
	}
}

func (user *User) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "User",
		Description: "A user",
		Fields: schema.GraphQLFieldSpecMap{
			"id": {
				Name:        "id",
				Description: "The id of user.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, user.Id, f)
				},
			},
			"todos": {
				Name:        "todos",
				Description: "The todos for a user.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					if f.Arguments[0].Name == "status" {
						if f.Arguments[0].Value.(*graphql.Value).Value == "completed" {
							return r.Resolve(ctx, user.CompletedTodos, f)
						}
						if f.Arguments[0].Value.(*graphql.Value).Value == "active" {
							return r.Resolve(ctx, user.ActiveTodos, f)
						}
					}
					return r.Resolve(ctx, user.AnyTodos, f)
				},
			},
			"completedCount": {
				Name:        "completedCount",
				Description: "The todos for a user.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, user.AnyTodos.CompletedCount, f)
				},
			},
			"totalCount": {
				Name:        "totalCount",
				Description: "The todos for a user.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, user.AnyTodos.TotalCount, f)
				},
			},
		},
	}
}
