package internal

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/graceful"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func SetupAPIServer() (context.Context, context.CancelFunc, *graceful.Graceful, error) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	r, err := graceful.Default()
	if err != nil {
		return ctx, stop, nil, err
	}
	setupRoutes(r)
	return ctx, stop, r, nil
}

func standardLogger() gin.HandlerFunc {
	return logger.SetLogger(
		logger.WithUTC(true),
		logger.WithLogger(func(c *gin.Context, l zerolog.Logger) zerolog.Logger {
			return l.Output(gin.DefaultWriter).
				With().
				Logger()
		}),
	)
}
