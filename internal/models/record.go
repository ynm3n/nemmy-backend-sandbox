package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Record struct {
	bun.BaseModel `bun:"table:records"`

	ID        int64     `json:"-" bun:",pk,autoincrement"`
	SubjectID int64     `json:"-" bun:",nullzero,notnull"`
	StartedAt time.Time `json:"startedAt" bun:",nullzero,notnull,default:current_timestamp"`
	Duration  int64     `json:"duration" bun:",notnull"`
	Comment   string    `json:"comment" bun:",type:text"`

	Subject Subject `json:"-" bun:"rel:belongs-to,join:subject_id=id"`
}
