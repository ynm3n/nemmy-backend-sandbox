package internal

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID    int64  `json:"userid" bun:"id,pk,nullzero,notnull"`
	Name  string `json:"username" bun:",nullzero,notnull"`
	Email string `json:"email" bun:",nullzero,notnull"`

	CreatedAt time.Time `json:"createdAt" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"updatedAt" bun:",nullzero,notnull,default:current_timestamp"`
}
