//go:build integration

package test

import (
	"context"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"user_base/config"
	"user_base/internal/app"
	"user_base/pkg/httpserver"
	"user_base/pkg/postgres"
	"user_base/pkg/profile_client"
)

// Prepare:  make up
// Run test: make integration-test

var ctx = context.Background()

func Test_Integration(t *testing.T) {
	suite.Run(t, &Suite{})
}

type Suite struct {
	suite.Suite
	*require.Assertions

	profile *profile_client.Client
}

func (s *Suite) SetupSuite() { // В начале всех тестов
	s.Assertions = s.Require()

	s.ResetMigrations()

	// Config
	c := config.Config{
		App: config.App{
			Name:    "user_base",
			Version: "v0.1.1",
		},
		Postgres: postgres.Config{
			User:     "login",
			Password: "pass",
			Port:     "5432",
			Host:     "localhost",
			DBName:   "postgres",
		},
		HTTP: httpserver.Config{
			Port:     "8080",
			BasePath: "/api/v1",
		},
	}

	// Logger and OTEL disable
	log.Logger = zerolog.Nop()

	// Server
	go func() {
		err := app.Run(context.Background(), c)
		s.NoError(err)
	}()

	// API client
	s.profile = profile_client.New(profile_client.Config{Host: "localhost", Port: "8080"})

	for i := 0; i < 20; i++ {
		_, err := s.profile.GetProfile(ctx, "health-check")
		if err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (s *Suite) TearDownSuite() {} // В конце всех тестов

func (s *Suite) SetupTest() {} // Перед каждым тестом

func (s *Suite) TearDownTest() {} // После каждого теста
