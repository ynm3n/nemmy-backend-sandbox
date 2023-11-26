package internal

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	fixturepkg "github.com/ynm3n/go-bun-exercise/internal/fixture"
)

// DBに接続し、指定したモデル(テーブル)を作り直し、フィクスチャ(ダミーデータ)を登録する
// マイグレーションが難しそうなので、開発ではひとまずこれを使ってください
func RecreateDB(ctx context.Context, dsn string, models []any) (*bun.DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	if err := sqldb.PingContext(ctx); err != nil {
		return nil, err
	}
	db := bun.NewDB(sqldb, pgdialect.New())
	db.RegisterModel(models...)
	fixture := dbfixture.New(db, dbfixture.WithRecreateTables())
	if err := fixture.Load(ctx, fixturepkg.EmbedFS, "fixture.yml"); err != nil {
		return nil, err
	}
	return db, nil
}

// 未実装です
// func NewDB(ctx context.Context, dsn string, models []any) (*bun.DB, error) {
// 	panic("unimplemented")
// 	return nil, nil
// }

func BuildDSN(cfg *Config) string {
	return fmt.Sprintf("postgres://%[1]v:%[2]v@db/%[1]v?sslmode=disable", cfg.DBUser, cfg.DBPassword)
}
