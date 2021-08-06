package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os/signal"
	"strings"
	"syscall"

	"github.com/libmonsoon-dev/web-crawler/storage/sql/postgres/migration"
)

func main() {
	cmd := flag.String("cmd", "up", "Command to execute. Available commands: 'up', 'down'")
	connection := flag.String("conn", "postgres://postgres:devpass@localhost:5432/postgres?sslmode=disable", "Database connection string")

	flag.Parse()

	db, err := sql.Open("postgres", *connection)
	check(err)

	migrator := migration.NewMigrator(db)
	fn := migrator.Up

	switch strings.ToLower(*cmd) {
	case "up":
		fn = migrator.Up
	case "down":
		fn = migrator.Down
	default:
		check(fmt.Errorf("invalid command %q", *cmd))
	}

	ctx, stopNotify := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer stopNotify()

	check(fn(ctx))

	fmt.Println("Done")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
