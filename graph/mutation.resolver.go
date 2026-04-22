package graph

import (
	"context"
	"example/graph/model"
	"fmt"
	"github.com/google/uuid"
)

// This is where Resolver is embedded as a pointer

type mutationResolver struct {
	*Resolver
}

// mutationResolver has Resolver embedded.
// *model.Todo these are usually resolvers because 
// importing from anywhere always starts with the package name
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	
	var createdByUser *model.User


	for _,user := range r.UserData {
		if  user.ID == input.UserID {
			createdByUser = user
			break
		}
	}

	if createdByUser == nil {
		return nil,fmt.Errorf("Could not find that user")
	}

	createdTodo := &model.Todo{
		ID:   uuid.NewString(),
		Text: input.Text,
		Done: false,
		User: createdByUser,
	}

	r.TodoData = append(r.TodoData, createdTodo)
	
	return createdTodo , nil
}