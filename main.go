package main

import (
	"github.com/jcheyer/jitsi-bg/internal/router"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("starting jitsi-bg...")

	r, err := router.New(
		router.WithListener(":8181"),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create router")
	}

	log.Fatal().Err(r.Run()).Send()
}
