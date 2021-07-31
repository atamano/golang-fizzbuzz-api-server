package main

import (
	"github.com/atamano/fizz-buzz/pkg/logger"
	"github.com/go-pg/migrations/v8"
)

const requestStatsTable = `
	CREATE TABLE fizzbuzz_requests_stats (
	id serial NOT NULL,
	key text NOT NULL UNIQUE,
	params json NOT NULL,
	counter int NOT NULL DEFAULT 0,
	PRIMARY KEY (id)
)`

func init() {
	up := []string{
		requestStatsTable,
	}

	down := []string{
		`DROP TABLE fizzbuzz_requests_stats`,
	}

	migrations.Register(func(db migrations.DB) error {
		logger.Info("creating initial tables")
		for _, q := range up {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		logger.Info("dropping initial tables")
		for _, q := range down {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
