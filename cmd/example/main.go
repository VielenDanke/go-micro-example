package main

import (
	"fmt"

	jsoncodec "github.com/unistack-org/micro-codec-json/v3"
	httpsrv "github.com/unistack-org/micro-server-http/v3"
	"github.com/unistack-org/micro/v3"
	"github.com/unistack-org/micro/v3/server"
	"github.com/vielendanke/go-micro-example/internal/example/handler"
	pb "github.com/vielendanke/go-micro-example/proto"
)

func main() {
	srv := micro.Server(httpsrv.NewServer(
		server.Address(":5050"),
		server.Codec("application/json", jsoncodec.NewCodec()),
	))

	svc := micro.NewService(srv)

	svc.Init()

	h := handler.NewMessageHandler()

	pb.RegisterPostServer(svc.Server(), h)

	if err := svc.Run(); err != nil {
		fmt.Println(err)
	}
}
