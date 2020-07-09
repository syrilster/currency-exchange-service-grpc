package exchange

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/syrilster/currency-exchange-service-grpc/internal/customhttp"
	"io/ioutil"
	"net/http"
)

type ClientInterface interface {
	GetExchangeRate(ctx context.Context, request Request) (*Response, error)
}

func NewClient(endpoint string, httpCommand customhttp.HTTPCommand, appID string) *client {
	return &client{
		URL:         endpoint,
		HttpCommand: httpCommand,
		AppID:       appID,
	}
}

type client struct {
	URL         string
	HttpCommand customhttp.HTTPCommand
	AppID       string
}

func (c *client) GetExchangeRate(ctx context.Context, request Request) (*Response, error) {
	contextLogger := log.WithContext(ctx)

	httpRequest, err := http.NewRequest(http.MethodGet, c.buildCurrencyExchangeEndpoint(request.FromCurrency, request.ToCurrency), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HttpCommand.Do(httpRequest)
	if err != nil {
		contextLogger.WithError(err).Errorf("there was an error calling the currency exchange API. %v", err)
		return nil, err
	}

	defer func() {
		if err = resp.Body.Close(); err != nil {
			fmt.Println("Error when closing:", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		contextLogger.Infof("status returned from currency exchange service %s", resp.Status)
		return nil, fmt.Errorf("currency exchange service returned status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		contextLogger.WithError(err).Errorf("error reading currency exchange service data resp body (%s)", err)
		return nil, err
	}

	response := &Response{}
	if err := json.Unmarshal(body, response); err != nil {
		contextLogger.WithError(err).Errorf("there was an error un marshalling the currency exchange API resp. %v", err)
		return nil, err
	}

	return response, nil
}

func (c *client) buildCurrencyExchangeEndpoint(from string, to string) (endpoint string) {
	return c.URL + "?app_id=" + c.AppID
}
