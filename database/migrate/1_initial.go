package main

import (
	"fmt"

	"github.com/go-pg/migrations"
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
		fmt.Println("creating initial tables")
		for _, q := range up {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("dropping initial tables")
		for _, q := range down {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
