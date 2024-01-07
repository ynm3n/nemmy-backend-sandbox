package main

import (
	"context"
	"log"
	"os"

	"github.com/ynm3n/go-bun-exercise/internal"
)

func main() {
	ctx := context.Background()
	cfg, err := internal.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := internal.RunApp(ctx, cfg); err != nil {
		panic(err) // TODO: エラー処理
		os.Exit(1)
	}
}
