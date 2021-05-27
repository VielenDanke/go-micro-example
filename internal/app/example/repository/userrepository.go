package repository

import (
	"context"

	"github.com/vielendanke/go-micro-example/internal/app/example/model"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
}
