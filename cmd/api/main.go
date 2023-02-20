package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/xkamail/dotconfig"
)

type Config struct {
	Port string `env:"PORT" default:"8080"`
}

var cfg Config

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
func run() error {
	if err := dotconfig.Load(&cfg, "./.env"); err != nil {
		return err
	}
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	return nil
}
