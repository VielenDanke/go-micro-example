package service

import (
	"context"

	pb "github.com/vielendanke/go-micro-example/proto"
)

type UserService interface {
	FindByID(ctx context.Context, id string) (*pb.UserModel, error)
	FindAll(ctx context.Context) ([]*pb.UserModel, error)
}
