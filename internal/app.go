package internal

import (
	"context"
	"fmt"

	"github.com/ynm3n/go-bun-exercise/internal/models"
)

func RunApp(ctx context.Context) error {
	cfg, err := GetConfig()
	if err != nil {
		return err
	}

	models := []any{ // 新しいテーブル(モデル)を作ったらここに書きましょう
		(*models.User)(nil),
		(*models.Subject)(nil),
		(*models.Record)(nil),
		// (*NewModelType)(nil),
	}
	db, err := RecreateDB(ctx, BuildDSN(cfg), models)
	if err != nil {
		return err
	}
	defer db.Close()

	e := NewEcho(cfg, db)

	addr := fmt.Sprintf("0.0.0.0:%v", cfg.Port)
	if err := e.Start(addr); err != nil {
		return err
	}

	return nil
}
