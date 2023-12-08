package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/uptrace/bun/migrate"
	"github.com/ynm3n/go-bun-exercise/internal/migrations"
)

var name string

const usage = "Migration name. Except timestamp because auto added."

func main() {
	flag.Usage = func() {
		fmt.Println("This is migration file creator")
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "name", "", usage)
	flag.Parse()

	m := migrate.NewMigrator(nil, migrations.Migrations)
	if _, err := m.CreateGoMigration(context.Background(), name); err != nil {
		panic(err)
	}
}
