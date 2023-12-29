package models

import "github.com/uptrace/bun"

type Subject struct {
	bun.BaseModel `bun:"table:subjects"`

	ID          int64  `json:"id" bun:",pk,autoincrement"`
	UserID      int64  `json:"-"`
	Name        string `json:"name" bun:",nullzero,notnull,type:varchar(255)"`
	Description string `json:"description" bun:",type:text"`

	Records []*Record `json:"-" bun:"rel:has-many,join:id=subject_id"`
	User    User      `json:"-" bun:"rel:belongs-to,join:user_id=id"`
}
