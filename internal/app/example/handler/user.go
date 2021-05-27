package handler

import (
	"context"
	"net/http"

	hsrv "github.com/unistack-org/micro-server-http/v3"
	"github.com/vielendanke/go-micro-example/internal/app/example/service"
	pb "github.com/vielendanke/go-micro-example/proto"
)

type UserHandler struct {
	s service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{s: s}
}

func (h *UserHandler) FindByID(ctx context.Context, req *pb.FindByIDRequest, rsp *pb.FindByIDResponse) error {
	usr, err := h.s.FindByID(ctx, req.UserId)
	if err != nil {
		return determineErrorStatus(ctx, err)
	}
	rsp.Username = usr.Username
	return nil
}

func determineErrorStatus(ctx context.Context, err error) error {
	_, isNotFound := err.(*pb.ErrorNotFound)
	_, isInternal := err.(*pb.Error)
	if isNotFound {
		hsrv.SetRspCode(ctx, http.StatusNotFound)
	}
	if isInternal {
		hsrv.SetRspCode(ctx, http.StatusInternalServerError)
	}
	return hsrv.SetError(err)
}
