package example

import (
	"context"
	"database/sql"

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
	"github.com/vielendanke/go-micro-example/internal/app/example/repository"
	"github.com/vielendanke/go-micro-example/internal/app/example/service"
	pb "github.com/vielendanke/go-micro-example/proto"
)

func initDB(name, url string) (*sql.DB, error) {
	db, openErr := sql.Open(name, url)

	if openErr != nil {
		return nil, openErr
	}
	if pingErr := db.Ping(); pingErr != nil {
		return nil, pingErr
	}
	return db, nil
}

func StartExampleService(ctx context.Context, errCh chan<- error) {
	cfg := configs.NewConfig()

	if loadErr := config.Load(ctx,
		config.NewConfig(
			config.Struct(cfg),
		),
		fileconfig.NewConfig(
			config.AllowFail(true),
			config.Struct(cfg),
			config.Codec(jsoncodec.NewCodec()),
			fileconfig.Path("./config.json"),
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
	db, dbErr := initDB(cfg.DB.Name, cfg.DB.URL)

	if dbErr != nil {
		errCh <- dbErr
		return
	}
	r := repository.NewUserRepository(db)

	s := service.NewUserService(r)

	h := handler.NewUserHandler(s)

	if err := pb.RegisterUserServer(svc.Server(), h); err != nil {
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
