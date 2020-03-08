package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/atamano/fizz-buzz/database"
	"github.com/go-pg/migrations"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
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
		logrus.WithError(err).Fatal("Failed to parse database config")
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
	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
