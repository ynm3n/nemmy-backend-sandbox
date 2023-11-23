package internal

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewDB(ctx context.Context, dsn string) (*bun.DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	if err := sqldb.PingContext(ctx); err != nil {
		return nil, err
	}
	db := bun.NewDB(sqldb, pgdialect.New())
	return db, nil
}

func BuildDSN(cfg *Config) string {
	return fmt.Sprintf("postgres://%[1]v:%[2]v@db/%[1]v?sslmode=disable", cfg.DBUser, cfg.DBPassword)
}
