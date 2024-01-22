package main

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/ymylei/playground-project/service-api/internal"
)

func main() {
	ctx, stop, r, err := internal.SetupAPIServer()
	defer stop()
	if err != nil {
		log.Fatal().Err(err).Msg("web server setup failure")
	}
	defer r.Close()
	err = r.RunWithContext(ctx)
	if err != nil && err != context.Canceled {
		log.Fatal().Err(err).Msg("web server crash")
	}
}
