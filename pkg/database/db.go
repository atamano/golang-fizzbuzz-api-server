package database

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	Debug    bool
}

type DB struct {
	db *pg.DB
}

type Query interface {
	orm.DB
	Begin() (*pg.Tx, error)
}

func checkConnection(db *pg.DB) error {
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

func Connect(config Config) (DB, error) {
	db := pg.Connect(&pg.Options{
		User:     config.User,
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		Database: config.Name,
	})

	if err := checkConnection(db); err != nil {
		return DB{}, err
	}

	if config.Debug {
		db.AddQueryHook(sqlHook{})
	}

	return DB{db}, nil
}

func (d DB) Close() {
	d.db.Close()
}

func (d DB) DB() Query {
	return d.db
}
