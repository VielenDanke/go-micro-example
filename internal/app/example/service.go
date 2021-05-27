package example

import (
	"context"

	hcli "github.com/unistack-org/micro-client-http/v3"
	jsoncodec "github.com/unistack-org/micro-codec-json/v3"
	envconfig "github.com/unistack-org/micro-config-env/v3"
	fileconfig "github.com/unistack-org/micro-config-file/v3"
	httpsrv "github.com/unistack-org/micro-server-http/v3"
	"github.com/unistack-org/micro/v3"
	"github.com/unistack-org/micro/v3/client"
	"github.com/unistack-org/micro/v3/config"
	"github.com/unistack-org/micro/v3/logger"
	"github.com/unistack-org/micro/v3/server"
	"github.com/vielendanke/go-micro-example/configs"
	"github.com/vielendanke/go-micro-example/internal/app/example/handler"
	pb "github.com/vielendanke/go-micro-example/proto"
)

func StartExampleService(ctx context.Context, errCh chan<- error) {
	cfg := configs.NewConfig()

	if loadErr := config.Load(ctx,
		config.NewConfig(
			config.Struct(cfg),
		),
		fileconfig.NewConfig(
			config.AllowFail(true),
			config.Struct(cfg),
			fileconfig.Path("./configs.json"),
		),
		envconfig.NewConfig(
			config.AllowFail(true),
			config.Struct(cfg),
		),
	); loadErr != nil {
		errCh <- loadErr
		return
	}
	svc := micro.NewService()

	if initErr := svc.Init(); initErr != nil {
		errCh <- initErr
		return
	}

	if initErr := svc.Init(
		micro.Server(httpsrv.NewServer(
			server.Address(cfg.Server.Addr),
			server.Name(cfg.Server.Name),
			server.Version(cfg.Server.Version),
			server.Codec("application/json", jsoncodec.NewCodec()),
		)),
		micro.Client(hcli.NewClient(
			client.ContentType("application/json"),
			client.Codec("application/json", jsoncodec.NewCodec()),
		)),
	); initErr != nil {
		errCh <- initErr
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
