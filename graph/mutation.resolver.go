package graph

import (
	"context"
	"example/graph/model"
	"fmt"
	"github.com/google/uuid"
)

type mutationResolver struct {
	*Resolver
}

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