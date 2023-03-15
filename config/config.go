package config

var (
	tgApiToken = "TELEGRAM_APITOKEN"
)

type Config struct {
	tgToken string
}

func BotConfig() *Config {
	return &Config{
		tgToken: tgApiToken,
	}
}
