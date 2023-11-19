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

	// db, err := NewDB(ctx, cfg)
	// if err != nil {
	// 	return err
	// }
	// defer db.Close()

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
