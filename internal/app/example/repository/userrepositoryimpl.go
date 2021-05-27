package repository

import (
	"context"
	"database/sql"

	"github.com/vielendanke/go-micro-example/internal/app/example/model"
	pb "github.com/vielendanke/go-micro-example/proto"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (u UserRepositoryImpl) FindByID(ctx context.Context, id string) (*model.User, error) {
	row := u.db.QueryRowContext(ctx, "select u.username from users u where u.id=$1", id)

	usr := &model.User{}

	scanErr := row.Scan(&usr.Username)

	if scanErr == sql.ErrNoRows {
		return nil, &pb.ErrorNotFound{Msg: scanErr.Error()}
	}
	if scanErr != nil {
		return nil, &pb.Error{Msg: scanErr.Error()}
	}
	return usr, nil
}
