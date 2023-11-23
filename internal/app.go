package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunApp(ctx context.Context) error {
	cfg, err := GetConfig()
	if err != nil {
		return err
	}

	models := []any{ // 新しいテーブル(モデル)を作ったらここに書きましょう
		(*User)(nil),
		// (*NewModelType)(nil),
	}
	db, err := NewDB(ctx, BuildDSN(cfg), models)
	if err != nil {
		return err
	}
	defer db.Close()

	e := echo.New()
	RegisterRoute(e)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{cfg.FrontendURL},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	addr := fmt.Sprintf("0.0.0.0:%v", cfg.Port)
	if err := e.Start(addr); err != nil {
		return err
	}

	return nil
}
