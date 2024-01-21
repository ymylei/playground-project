package main

import (
	"github.com/rs/zerolog/log"
	"github.com/ymylei/playground-project/service-api/internal"
)

func main() {
	err := internal.APIServer()
	if err != nil {
		log.Fatal().Err(err).Msg("web server crash")
	}
}
