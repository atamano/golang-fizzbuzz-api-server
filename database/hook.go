package database

import (
	"log"

	"github.com/go-pg/pg"
)

type sqlHook struct{}

func (l *sqlHook) BeforeQuery(e *pg.QueryEvent) {}

func (l *sqlHook) AfterQuery(e *pg.QueryEvent) {
	query, err := e.FormattedQuery()
	if err != nil {
		panic(err)
	}
	log.Println(query)
}
