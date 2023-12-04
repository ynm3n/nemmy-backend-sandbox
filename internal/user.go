package internal

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID              int64  `json:"-" bun:",pk,autoincrement"`
	Username        string `json:"username" bun:",unique,nullzero,notnull,type:varchar(255)"`
	DisplayName     string `json:"displayName" bun:",nullzero,notnull,type:varchar(255)"`
	ProfileImageURL string `json:"profileImageUrl" bun:",type:varchar(255)"`

	CreatedAt time.Time `json:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"-" bun:",nullzero,notnull,default:current_timestamp"`

	Subjects []*Subject `json:"-" bun:"rel:has-many,join:id=user_id"`
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
