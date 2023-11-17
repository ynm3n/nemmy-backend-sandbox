package main

import (
	"context"
	"os"

	"github.com/ynm3n/go-bun-exercise/internal"
)

func main() {
	ctx := context.Background()
	if err := internal.RunApp(ctx); err != nil {
		panic(err) // TODO: エラー処理
		os.Exit(1)
	}
}
