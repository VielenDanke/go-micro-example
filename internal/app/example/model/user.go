package model

import pb "github.com/vielendanke/go-micro-example/proto"

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

func MapToUserModel(u *User) *pb.UserModel {
	return &pb.UserModel{
		Id:       u.ID,
		Username: u.Username,
	}
}
