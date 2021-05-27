package handler

import (
	"context"
	"io/ioutil"
	"net/http"

	hsrv "github.com/unistack-org/micro-server-http/v3"
	"github.com/unistack-org/micro/v3/codec"
	"github.com/unistack-org/micro/v3/logger"
	"github.com/unistack-org/micro/v3/metadata"
	"github.com/vielendanke/go-micro-example/internal/app/example/service"
	pb "github.com/vielendanke/go-micro-example/proto"
)

var (
	applicationJson = "application/json"
	applicationXml  = "application/xml"
)

type UserHandler struct {
	s service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{s: s}
}

func (h *UserHandler) FindAll(ctx context.Context, req *pb.FindAllRequest, rsp *pb.FindAllResponse) error {
	users, err := h.s.FindAll(ctx)

	if err != nil {
		return determineErrorStatus(ctx, err)
	}
	rsp.Users = users
	determineResponseContentType(ctx)
	return nil
}

func (h *UserHandler) FindByID(ctx context.Context, req *pb.FindByIDRequest, rsp *pb.FindByIDResponse) error {
	usr, err := h.s.FindByID(ctx, req.UserId)

	if err != nil {
		return determineErrorStatus(ctx, err)
	}
	rsp.User = usr
	determineResponseContentType(ctx)
	return nil
}

func (h *UserHandler) DownloadUserFile(ctx context.Context, req *pb.DownloadRequest, rsp *codec.Frame) error {
	_, uErr := h.s.FindByID(ctx, req.UserId)

	if uErr != nil {
		return determineErrorStatus(ctx, uErr)
	}
	data, err := ioutil.ReadFile("build.sh")

	if err != nil {
		return determineErrorStatus(ctx, &pb.Error{Msg: err.Error()})
	}
	rsp.Data = data
	return nil
}

func determineResponseContentType(ctx context.Context) {
	m, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		logger.Warn(ctx, "Metadata in context not found")
		return
	}
	ah, _ := m.Get("Accept")

	switch ah {
	case applicationJson:
		m.Set("Content-Type", applicationJson)
	case applicationXml:
		m.Set("Content-Type", applicationXml)
	default:
		logger.Info(ctx, "No accept content type, use default application/json")
		return
	}
	if isSet := metadata.SetOutgoingContext(ctx, m); !isSet {
		logger.Warn(ctx, "Metadata is not being set")
		return
	}
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
