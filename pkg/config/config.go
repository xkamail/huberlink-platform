package config

type Config struct {
	Port                string `env:"PORT"`
	DatabaseURL         string `env:"DATABASE_URL"`
	JWTSecret           string `env:"JWT_SECRET"`
	DiscordClientID     string `env:"DISCORD_CLIENT_ID"`
	DiscordClientSecret string `env:"DISCORD_CLIENT_SECRET"`
	DiscordRedirectURI  string `env:"DISCORD_CLIENT_REDIRECT_URI"`
}

var cfg Config

func Init() {

}

func Load() *Config {
	return &cfg
}
