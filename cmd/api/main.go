package main

import (
	"context"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/moonrhythm/parapet"
	"github.com/moonrhythm/parapet/pkg/cors"

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
	cfg := config.Load()
	conn, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
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
	srv.Addr = net.JoinHostPort("", cfg.Port)
	log.Printf("serve http api on %s\n", srv.Addr)
	return srv.ListenAndServe()
}
