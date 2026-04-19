package main

import (
	"context"
	"user_base/config"
	"user_base/internal/app"
	"user_base/pkg/logger"

	"github.com/rs/zerolog/log"
)

func main() {
	c, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("config.New")
	}

	ctx := context.Background()

	logger.Init(c.Logger)

	err = app.Run(ctx, c)
	if err != nil {
		log.Error().Err(err).Msg("app.Run")
	}
}
