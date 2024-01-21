package internal

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/graceful"
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

	// End Routes Here
	err = r.RunWithContext(ctx)
	if err != nil {
		return err
	}
}
