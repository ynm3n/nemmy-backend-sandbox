package internal

import (
	"context"
	"fmt"
)

func RunApp(ctx context.Context, cfg *Config) error {
	db, err := NewDB(ctx, BuildDSN(cfg))
	if err != nil {
		return fmt.Errorf("RunApp: %w", err)
	}
	defer db.Close()
	if err := Migrate(ctx, db); err != nil {
		return fmt.Errorf("RunApp: %w", err)
	}

	e := NewEcho(cfg, db)

	addr := fmt.Sprintf("0.0.0.0:%v", cfg.Port)
	if err := e.Start(addr); err != nil {
		return fmt.Errorf("RunApp: %w", err)
	}

	return nil
}
