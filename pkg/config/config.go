package config

import (
	"github.com/xkamail/dotconfig"
)

type Config struct {
	Port                string `env:"PORT" required:"true"`
	DatabaseURL         string `env:"DATABASE_URL"`
	JWTSecret           string `env:"JWT_SECRET"`
	DiscordClientID     string `env:"DISCORD_CLIENT_ID"`
	DiscordClientSecret string `env:"DISCORD_CLIENT_SECRET"`
	DiscordRedirectURI  string `env:"DISCORD_CLIENT_REDIRECT_URI"`
}

var cfg Config

func Init() error {
	if err := dotconfig.Load(&cfg, ".env"); err != nil {
		return err
	}
	return nil
}

func Load() *Config {
	return &cfg
}
