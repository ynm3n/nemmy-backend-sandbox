package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"
)

func NewEcho(cfg *Config, db *bun.DB) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{cfg.FrontendURL},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	h := &Handler{DB: db}
	api := e.Group("/api")

	u := api.Group("/users")
	u.GET("/:username", h.GetUser)
	return e
}

type Handler struct {
	DB *bun.DB
}
