package server

import (
	"context"

	jsoncodec "github.com/unistack-org/micro-codec-json/v3"
	httpsrv "github.com/unistack-org/micro-server-http/v3"
	"github.com/unistack-org/micro/v3"
	"github.com/unistack-org/micro/v3/logger"
	"github.com/unistack-org/micro/v3/server"
	"github.com/vielendanke/go-micro-example/internal/example/handler"
	pb "github.com/vielendanke/go-micro-example/proto"
)

func StartExampleService(ctx context.Context, errCh chan<- error) {
	srv := micro.Server(httpsrv.NewServer(
		server.Address(":5050"),
		server.Codec("application/json", jsoncodec.NewCodec()),
	))

	svc := micro.NewService(srv)

	if err := svc.Init(); err != nil {
		logger.Errorf(ctx, "Error initializing service %v", err)
		errCh <- err
		return
	}

	h := handler.NewMessageHandler()

	if err := pb.RegisterPostServer(svc.Server(), h); err != nil {
		logger.Errorf(ctx, "Error registering server %v", err)
		errCh <- err
		return
	}

	if err := svc.Run(); err != nil {
		logger.Errorf(ctx, "Error runnig service %v", err)
		errCh <- err
		return
	}
}
