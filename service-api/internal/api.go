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

func APIServer() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r, err := graceful.Default()
	if err != nil {
		return err
	}
	defer r.Close()

	// Insert Routes Here
	r.GET("/ready", standardLogger(), healthCheck)

	// End Routes Here
	err = r.RunWithContext(ctx)
	if err != nil && err != context.Canceled {
		return err
	}
	return nil
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
