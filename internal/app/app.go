package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"user_base/pkg/router"

	"user_base/config"
	"user_base/internal/usecase"
	"user_base/pkg/httpserver"
	pgpool "user_base/pkg/postgres"
	"user_base/pkg/transaction"

	"user_base/internal/adapter/postgres"

	"github.com/rs/zerolog/log"
)

func Run(ctx context.Context, c config.Config) error {
	pgPool, err := pgpool.New(ctx, c.Postgres)
	if err != nil {
		return fmt.Errorf("postgres.New: %w", err)
	}

	transaction.Init(pgPool.Pool)

	uc := usecase.New(postgres.New())

	r := router.New(c.HTTP.BasePath, uc)

	httpServer := httpserver.New(r, c.HTTP)

	log.Info().Msg("app: started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Info().Msg("app: got signal to stop")

	pgPool.Close()
	httpServer.Close()

	log.Info().Msg("app: stopped")

	return nil
}
