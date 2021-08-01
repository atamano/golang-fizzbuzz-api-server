package database

import (
	"context"

	"github.com/atamano/fizz-buzz/pkg/logger"
	"github.com/go-pg/pg/v10"
)

type sqlHook struct{}

func (l sqlHook) BeforeQuery(ctx context.Context, e *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (l sqlHook) AfterQuery(ctx context.Context, e *pg.QueryEvent) error {
	query, err := e.FormattedQuery()

	logger.Info(query)

	return err
}
