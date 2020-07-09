package main

import (
	"github.com/syrilster/currency-exchange-service-grpc/internal"
	"github.com/syrilster/currency-exchange-service-grpc/internal/config"
)

func main() {
	cfg := config.NewApplicationConfig()
	server := internal.SetupServer(cfg)
	server.Start("", cfg.ServerPort())
}
