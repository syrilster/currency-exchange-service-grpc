package config

import (
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

type envConfig struct {
	LogLevel   string `env:"LOG_LEVEL"`
	ServerPort int    `env:"SERVER_PORT" envDefault:"8000"`
	AppID      string `env:"APP_ID" envDefault:"76d9797902004813a592f6b2402590a9"`

	CurrencyExchangeEndpoint string `env:"CURRENCY_EXCHANGE_ENDPOINT" envDefault:"https://openexchangerates.org/api/latest.json"`
}

func newEnvironmentConfig() *envConfig {
	cfg := &envConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal("cannot find configs for server: \n", err)
	}
	return cfg
}
