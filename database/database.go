package database

import (
	"fmt"

	"github.com/go-pg/pg"
)

//Config for database
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	Debug    bool
}

func checkConnection(db *pg.DB) error {
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	return err
}

// Connect returns a postgres connection pool.
func Connect(config Config) (*pg.DB, error) {

	db := pg.Connect(&pg.Options{
		User:     config.User,
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		Database: config.Name,
	})

	if err := checkConnection(db); err != nil {
		return nil, err
	}

	if config.Debug {
		db.AddQueryHook(&sqlHook{})
	}

	return db, nil
}
