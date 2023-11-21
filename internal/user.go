package internal

import (
	"time"

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
