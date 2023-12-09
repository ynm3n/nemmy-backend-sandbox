package models

import (
	"time"

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
