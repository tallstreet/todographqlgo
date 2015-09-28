package graph

import (
	"github.com/tallstreet/graphql"
	"github.com/tallstreet/graphql/executor/resolver"
	"github.com/tallstreet/graphql/schema"
	"golang.org/x/net/context"
)

type User struct {
	Id    string
	Todos *TodoConnection
}

func (user *User) addToDo(todo *TodoNode) {
	user.Todos.addTodo(todo)
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
					return r.Resolve(ctx, user.Todos, f)
				},
			},
			"completedCount": {
				Name:        "completedCount",
				Description: "The todos for a user.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, user.Todos.CompletedCount, f)
				},
			},
			"totalCount": {
				Name:        "totalCount",
				Description: "The todos for a user.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, user.Todos.TotalCount, f)
				},
			},
		},
	}
}
