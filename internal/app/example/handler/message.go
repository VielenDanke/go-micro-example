package handler

import (
	"context"

	"github.com/google/uuid"
	httpsrv "github.com/unistack-org/micro-server-http/v3"
	pb "github.com/vielendanke/go-micro-example/proto"
)

type MessageHandler struct {
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

func (h *MessageHandler) FindByID(ctx context.Context, req *pb.FindByIDRequest, rsp *pb.FindByIDResponse) error {
	rsp.PostId = uuid.New().String()
	rsp.Text = "Hello"
	httpsrv.SetRspCode(ctx, 200)
	return nil
}
