package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/moonrhythm/parapet"
	"github.com/moonrhythm/parapet/pkg/cors"
	"golang.org/x/exp/slog"

	"github.com/xkamail/huberlink-platform/iot"
	"github.com/xkamail/huberlink-platform/pkg/config"
	"github.com/xkamail/huberlink-platform/pkg/pgctx"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	if err := config.Init(); err != nil {
		return err
	}

	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer conn.Close()
	if err := conn.Ping(context.Background()); err != nil {
		return err
	}
	srv := parapet.NewBackend()
	srv.Use(parapet.MiddlewareFunc(pgctx.Middleware(conn)))
	srv.Use(cors.New())
	srv.Handler = iot.Handlers()
	srv.Addr = net.JoinHostPort("", config.Load().Port)
	slog.Info("serving http api on ", "addr", srv.Addr)
	return srv.ListenAndServe()
}
