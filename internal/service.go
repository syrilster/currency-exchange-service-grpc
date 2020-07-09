package internal

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/syrilster/currency-exchange-service-grpc/gRpc/gen"
	"github.com/syrilster/currency-exchange-service-grpc/internal/exchange"
	"strings"
)

type Service struct {
	client exchange.ClientInterface
}

func NewService(c exchange.ClientInterface) *Service {
	return &Service{client: c}
}

func (s *Service) GetExchangeRate(ctx context.Context, req *gen.Request) (*gen.Response, error) {
	ctxLogger := log.WithContext(ctx)
	ctxLogger.Infof("Calling the exchange (openexchangerates) fetch API")

	request := toExchangeRequest(req)
	resp, err := s.client.GetExchangeRate(ctx, request)
	if err != nil {
		ctxLogger.Infof("Failed to fetch currency exchange rate: %v", err)
		return nil, err
	}
	response := unMarshallExchangeRate(resp, req)
	return response, nil
}

func unMarshallExchangeRate(resp *exchange.Response, req *gen.Request) *gen.Response {
	var fromCurrency = req.FromCurrency
	var toCurrency = req.ToCurrency
	var conversionMultiple float32
	var exchangeRate float32
	if strings.EqualFold(fromCurrency, "USD") {
		exchangeRate = getRateForCurrency(resp.Rates, toCurrency)
		conversionMultiple = exchangeRate
	} else if strings.EqualFold(toCurrency, "USD") {
		exchangeRate = getRateForCurrency(resp.Rates, fromCurrency)
		conversionMultiple = float32(1) / exchangeRate
	} else {
		// FromCurrency to USD and then USD to toCurrency
		exchangeRate = getRateForCurrency(resp.Rates, toCurrency)
		usdToFromCurrency := getRateForCurrency(resp.Rates, fromCurrency)
		toCurrencyToUSD := float32(1) / exchangeRate
		foreignCurrencyFactor := float32(1) / usdToFromCurrency
		conversionMultiple = foreignCurrencyFactor / toCurrencyToUSD
	}

	return &gen.Response{
		ConversionMultiple: conversionMultiple,
	}
}

func getRateForCurrency(rates map[string]interface{}, currency string) float32 {
	var exchangeRate float64
	for key, rate := range rates {
		if strings.EqualFold(key, currency) {
			exchangeRate = rate.(float64)
			break
		}
	}
	return float32(exchangeRate)
}

func toExchangeRequest(request *gen.Request) exchange.Request {
	return exchange.Request{
		FromCurrency: request.FromCurrency,
		ToCurrency:   request.ToCurrency,
	}
}
