package service

import (
	"context"

	"github.com/vielendanke/go-micro-example/internal/app/example/model"
	"github.com/vielendanke/go-micro-example/internal/app/example/repository"
)

type UserServiceImpl struct {
	r repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &UserServiceImpl{r: r}
}

func (s UserServiceImpl) FindByID(ctx context.Context, id string) (*model.User, error) {
	return s.r.FindByID(ctx, id)
}
