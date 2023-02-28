package config

import (
	"github.com/xkamail/dotconfig"
)

type Config struct {
	Port                string `envconfig:"PORT" required:"true"`
	DatabaseURL         string `envconfig:"DATABASE_URL" required:"true"`
	JWTSecret           string `envconfig:"JWT_SECRET" required:"true"`
	DiscordClientID     string `envconfig:"DISCORD_CLIENT_ID"`
	DiscordClientSecret string `envconfig:"DISCORD_CLIENT_SECRET"`
	DiscordRedirectUri  string `envconfig:"DISCORD_CLIENT_REDIRECT_URI" required:"true"`
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
