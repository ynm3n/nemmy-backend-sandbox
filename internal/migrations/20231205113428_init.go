package migrations

import (
	"context"

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
			ForeignKey(`("user_id") REFERENCES "users" ("id") ON DELETE CASCADE`).
			Exec(ctx); err != nil {
			return err
		}
		if _, err := db.NewCreateTable().
			Model((*models.Record)(nil)).
			ForeignKey(`("subject_id") REFERENCES "subjects" ("id") ON DELETE CASCADE`).
			Exec(ctx); err != nil {
			return err
		}
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
