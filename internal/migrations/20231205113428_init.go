package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/ynm3n/go-bun-exercise/internal/models"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		if _, err := db.NewCreateTable().
			Model((*models.User)(nil)).
			Exec(ctx); err != nil {
			return err
		}
		if _, err := db.NewCreateTable().
			Model((*models.Subject)(nil)).
			WithForeignKeys().
			Exec(ctx); err != nil {
			return err
		}
		if _, err := db.NewCreateTable().
			Model((*models.Record)(nil)).
			WithForeignKeys().
			Exec(ctx); err != nil {
			return err
		}
		fmt.Println("マイグレーション完了")
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		if _, err := db.NewDropTable().
			Model((*models.Record)(nil)).
			Exec(ctx); err != nil {
			return err
		}
		if _, err := db.NewDropTable().
			Model((*models.Subject)(nil)).
			Exec(ctx); err != nil {
			return err
		}
		if _, err := db.NewDropTable().
			Model((*models.User)(nil)).
			Exec(ctx); err != nil {
			return err
		}
		return nil
	})
}
