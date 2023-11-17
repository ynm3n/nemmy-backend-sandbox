package internal

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"` // デフォルトでも同じテーブル名になりそうなのでこのタグは不要かも

	ID    string `bun:",pk,nullzero,notnull"`
	Name  string `bun:",nullzero,notnull"`
	Email string `bun:",nullzero,notnull"`
	Age   int
}
