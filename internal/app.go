package internal

import (
	"context"
	"fmt"
)

func RunApp(ctx context.Context) error {
	cfg, err := GetConfig()
	if err != nil {
		return err
	}

	db, err := NewDB(ctx, BuildDSN(cfg))
	if err != nil {
		return err
	}
	defer db.Close()
	if err := Migrate(ctx, db); err != nil {
		return err
	}

	e := NewEcho(cfg, db)

	addr := fmt.Sprintf("0.0.0.0:%v", cfg.Port)
	if err := e.Start(addr); err != nil {
		return err
	}

	return nil
}
