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

	db, err := NewDB(ctx, cfg)
	// defer db.Close()
	if err != nil {
		return err
	}

	r, err := NewRouter(ctx, db)
	if err != nil {
		return err
	}

	p := fmt.Sprintf("0.0.0.0:%v", cfg.Port)
	if err := r.Start(p); err != nil {
		return err
	}

	return nil
}
