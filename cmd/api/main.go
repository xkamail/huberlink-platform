package main

import (
	"log"

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
	return nil
}
