package main

import (
	"context"

	"github.com/ynm3n/go-bun-exercise/internal"
)

func main() {
	ctx := context.Background()

	cfg, err := internal.GetConfig()
	if err != nil {
		panic(err)
	}

	db, err := internal.NewDB(ctx, cfg)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	if _, err := db.NewCreateTable().
		Model((*internal.User)(nil)).
		IfNotExists().
		Exec(ctx); err != nil {
		panic(err)
	}

}
