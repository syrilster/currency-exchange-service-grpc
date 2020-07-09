package internal

import (
	"context"
	"fmt"
	"github.com/syrilster/currency-exchange-service-grpc/gRpc/gen"
	"github.com/syrilster/currency-exchange-service-grpc/internal/exchange"
)

type Service struct {
	client exchange.ClientInterface
}

func NewService(c exchange.ClientInterface) *Service {
	return &Service{client: c}
}

func (s *Service) GetExchangeRate(ctx context.Context, req *gen.Request) (*gen.Response, error) {
	fmt.Println("Request: ", req.FromCurrency)
	return &gen.Response{ConversionMultiple: 50}, nil
}
