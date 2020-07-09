package internal

import (
	"fmt"
	"github.com/syrilster/currency-exchange-service-grpc/gRpc/gen"
	"github.com/syrilster/currency-exchange-service-grpc/internal/exchange"
	"google.golang.org/grpc"
	"log"
	"net"
)

// gRPC defines the grpc server struct
type gRPC struct {
	server *grpc.Server
}

type ServerConfig interface {
	CurrencyExchangeClient() exchange.ClientInterface
}

func SetupServer(cfg ServerConfig) *gRPC {
	server := grpc.NewServer()
	svc := NewService(cfg.CurrencyExchangeClient())
	gen.RegisterCurrencyExchangeServiceServer(server, svc)

	return &gRPC{server: server}
}

func (g gRPC) Start(addr string, port int) {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%v", addr, port))
	if err != nil {
		log.Fatal(err)
	}

	// Start gRPC server
	if err := g.server.Serve(l); err != nil {
		log.Fatal(err)
	}
}
