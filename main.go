package main

import (
	"context"
	"embed"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/unistack-org/micro/v3/logger"
	apiserver "github.com/vielendanke/go-micro-example/internal/app/example"
)

// Here we are using feature of GO called embed
// it is create embed file system inside, and insert files which are written below
// You can use it, but you have to write your parses of json file instead of using fileconfig.Path (see: service.go)
// We are often use this feature because all what we need to put in the container is our go binary
//go:embed config.json
var embedFS embed.FS

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	errCh := make(chan error, 1)

	logger.DefaultLogger = logger.NewLogger(logger.WithLevel(logger.TraceLevel))

	go apiserver.StartExampleService(ctx, embedFS, errCh)

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		errCh <- fmt.Errorf("%s", <-ch)
	}()

	logger.Infof(ctx, "Service terminated %v", <-errCh)
}
