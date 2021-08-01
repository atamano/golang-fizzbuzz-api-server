package main

import (
	"github.com/atamano/fizz-buzz/internal/fizzbuzz"
	"github.com/atamano/fizz-buzz/internal/statistics"
	"github.com/atamano/fizz-buzz/pkg/database"
	"github.com/atamano/fizz-buzz/pkg/logger"
	"github.com/atamano/fizz-buzz/pkg/server"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	server   server.Config
	database database.Config
}

func newConfig() config {
	var c config

	if err := envconfig.Process("server", &c.server); err != nil {
		logger.Fatal(err.Error(), "Failed to parse server config")
	}
	if err := envconfig.Process("database", &c.database); err != nil {
		logger.Fatal(err.Error(), "Failed to parse database config")
	}
	return c
}

func main() {
	config := newConfig()
	server := server.New(config.server)

	db, err := database.Connect(config.database)
	if err != nil {
		logger.Fatal(err.Error(), "Failed to connect database")
	}
	defer db.Close()

	fizzbuzzService := fizzbuzz.NewService()

	statsRepository := statistics.NewRepository(db.DB())
	statsService := statistics.NewService(statsRepository)

	v1 := server.NewGroup("/v1")
	fizzbuzz.RegisterHandlers(v1, fizzbuzzService, statsService)
	statistics.RegisterHandlers(v1, statsService)

	server.Run()
}
