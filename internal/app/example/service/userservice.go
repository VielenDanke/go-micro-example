package service

import (
	"context"

	"github.com/vielendanke/go-micro-example/internal/app/example/model"
)

type UserService interface {
	FindByID(ctx context.Context, id string) (*model.User, error)
}
