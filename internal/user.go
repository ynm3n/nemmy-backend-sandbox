package internal

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID          int64  `bun:",pk,autoincrement"`
	Username    string `json:"username" bun:",unique,nullzero,notnull"`
	DisplayName string `json:"displayName" bun:",nullzero,notnull"`
	Email       string `json:"email" bun:",unique,nullzero,notnull"`

	CreatedAt time.Time `json:"createdAt" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" bun:",nullzero,notnull,default:current_timestamp"`
}

func (h *Handler) GetUser(c echo.Context) error {
	p := c.Param("username")

	var u User
	err := h.DB.NewSelect().
		Model(&u).
		Where("username = ?", p).
		Scan(c.Request().Context())
	if err != nil {
		return err
	}

	if err := c.JSON(http.StatusOK, u); err != nil {
		return err
	}
	return nil
}
