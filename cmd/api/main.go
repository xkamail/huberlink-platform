package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/moonrhythm/parapet"
	"github.com/moonrhythm/parapet/pkg/cors"
	"github.com/xkamail/dotconfig"
	"golang.org/x/exp/slog"

	"github.com/xkamail/huberlink-platform/iot"
	"github.com/xkamail/huberlink-platform/pkg/pgctx"
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
	ctx := context.Background()
	if err := dotconfig.Load(&cfg, "./.env"); err != nil {
		return err
	}
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer conn.Close()
	pgctx.NewContext(ctx, conn)

	srv := parapet.NewBackend()
	srv.Use(parapet.MiddlewareFunc(pgctx.Middleware(conn)))
	srv.Use(cors.New())
	srv.Handler = iot.Handlers()
	slog.Info("serving http api on ", "addr", srv.Addr)
	return srv.ListenAndServe()
}
