package main

import (
	"flag"
	"os"

	"github.com/atamano/fizz-buzz/pkg/database"
	"github.com/atamano/fizz-buzz/pkg/logger"
	"github.com/go-pg/migrations/v8"
	"github.com/kelseyhightower/envconfig"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

func newConfig() database.Config {
	var c database.Config

	if err := envconfig.Process("database", &c); err != nil {
		logger.Fatal(err.Error(), "Failed to parse database config")
	}
	return c
}

func main() {
	flag.Usage = usage
	flag.Parse()

	config := newConfig()

	db, err := database.Connect(config)
	if err != nil {
		exitf(err.Error())
	}
	oldVersion, newVersion, err := migrations.Run(db.DB(), flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		logger.Infof("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		logger.Infof("version is %d\n", oldVersion)
	}
}

func usage() {
	logger.Info(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func exitf(s string, args ...interface{}) {
	logger.Errorf(s, args...)
	os.Exit(1)
}
