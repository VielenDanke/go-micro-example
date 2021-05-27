package middleware

import (
	"context"

	"github.com/unistack-org/micro/v3/logger"
	"github.com/unistack-org/micro/v3/metadata"
	"github.com/unistack-org/micro/v3/server"
)

func CORSMiddleware(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		err := fn(ctx, req, rsp)
		if err != nil {
			logger.Info(ctx, err.Error())
		}
		m, ok := metadata.FromOutgoingContext(ctx)

		if !ok {
			m = metadata.New(20)
		}
		setCORS(m)

		if isSet := metadata.SetOutgoingContext(ctx, m); !isSet {
			logger.Warn(ctx, "Context is not being set")
		}
		return err
	}
}

func setCORS(m metadata.Metadata) {
	m.Set("Access-Control-Allow-Origin", "*")
	m.Set("Access-Control-Allow-Method", "PUT, POST, HEAD, DELETE, GET, OPTIONS")
	m.Set("Access-Control-Allow-Headers", "*")
	m.Set("Access-Control-Max-Age", "86400")
	m.Set("Access-Control-Allow-Credentials", "true")
	m.Set("Access-Control-Expose-Headers", "Content-Type")
	m.Set("Vary", "Origin")
}
