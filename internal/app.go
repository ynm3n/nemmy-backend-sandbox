package internal

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
)

func RunApp(ctx context.Context) error {
	cfg, err := GetConfig()
	if err != nil {
		return err
	}

	// db, err := NewDB(ctx, cfg)
	// if err != nil {
	// 	return err
	// }
	// defer db.Close()

	e := echo.New()
	RegisterRoute(e)
	addr := fmt.Sprintf("0.0.0.0:%v", cfg.Port)
	if err := e.Start(addr); err != nil {
		return err
	}

	return nil
}
