package service

import (
	"context"

	"github.com/vielendanke/go-micro-example/internal/app/example/model"
	"github.com/vielendanke/go-micro-example/internal/app/example/repository"
	pb "github.com/vielendanke/go-micro-example/proto"
)

type UserServiceImpl struct {
	r repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &UserServiceImpl{r: r}
}

func (s UserServiceImpl) FindByID(ctx context.Context, id string) (*pb.UserModel, error) {
	u, err := s.r.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}
	return model.MapToUserModel(u), nil
}

func (s UserServiceImpl) FindAll(ctx context.Context) ([]*pb.UserModel, error) {
	users, err := s.r.FindAll(ctx)

	if err != nil {
		return nil, err
	}
	uModels := make([]*pb.UserModel, 0)

	for _, v := range users {
		uModels = append(uModels, model.MapToUserModel(v))
	}
	return uModels, nil
}
