package main

import (
	"github.com/atamano/fizz-buzz/database"
	"github.com/atamano/fizz-buzz/internal/fizzbuzz"
	"github.com/atamano/fizz-buzz/internal/statistics"
	"github.com/atamano/fizz-buzz/pkg/server"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

// Config of server
type Config struct {
	server   server.Config
	database database.Config
}

func newConfig() Config {
	var c Config

	if err := envconfig.Process("server", &c.server); err != nil {
		logrus.WithError(err).Fatal("Failed to parse server config")
	}
	if err := envconfig.Process("database", &c.database); err != nil {
		logrus.WithError(err).Fatal("Failed to parse database config")
	}
	return c
}

func main() {
	config := newConfig()
	server := server.New(config.server)

	db, err := database.Connect(config.database)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect database")
	}
	defer db.Close()

	// On large apps https://github.com/google/wire could be used to handle dependency injection
	fizzbuzzService := fizzbuzz.NewService()

	statsRepository := statistics.NewRepository(db)
	statsService := statistics.NewService(statsRepository)

	v1 := server.NewGroup("/v1")
	fizzbuzz.RegisterHandlers(v1, fizzbuzzService, statsService)
	statistics.RegisterHandlers(v1, statsService)

	server.Run()
}
