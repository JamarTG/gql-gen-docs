package graph

import (
	"context"
	"example/graph/model"
)

type queryResolver struct {
	Resolver *Resolver
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.Resolver.TodoData, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.Resolver.UserData, nil
}

func (r *userResolver) Friends(ctx context.Context, u *model.User) ([]*model.User, error) {
	return u.Friends,nil
}


func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

func (r *Resolver) User() UserResolver { return &userResolver{r} }