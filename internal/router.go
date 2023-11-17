package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func NewRouter(ctx context.Context, db *bun.DB) (*echo.Echo, error) {
	e := echo.New()
	e.GET("/api/users/:userID", func(c echo.Context) error {
		id := c.Param("userID")
		return c.String(http.StatusOK, fmt.Sprintf("hello, %v", id))
	})
	return e, nil
}
