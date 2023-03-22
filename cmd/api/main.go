package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/moonrhythm/parapet"
	"github.com/moonrhythm/parapet/pkg/cors"
	"golang.org/x/exp/slog"

	"github.com/xkamail/huberlink-platform/iot"
	"github.com/xkamail/huberlink-platform/pkg/api"
	"github.com/xkamail/huberlink-platform/pkg/config"
	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

var debug = flag.Bool("debug", false, "enable debug log")

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	l := slog.NewTextHandler(os.Stdout)
	if *debug {
		l.Enabled(context.TODO(), slog.LevelDebug)
	}
	slog.SetDefault(slog.New(l))
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
	srv.Use(parapet.MiddlewareFunc(recovery))
	srv.Handler = iot.Handlers()
	srv.Addr = net.JoinHostPort("", cfg.Port)
	log.Printf("serve http api on %s\n", srv.Addr)
	return srv.ListenAndServe()
}

// pretty print stack trace when program panic crash
func recovery(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				if rvr == http.ErrAbortHandler {
					// we don't recover http.ErrAbortHandler so the response
					// to the client is aborted, this should not be logged
					panic(rvr)
				}

				middleware.PrintPrettyStack(rvr)

				api.WriteError(w, uierr.InternalServer())
			}
		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
