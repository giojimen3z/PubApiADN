package client

import (
	"fmt"

	"github.com/PubApiADN/cmd/api/app/domain/model"
	"github.com/PubApiADN/cmd/api/app/infrastructure/config"
	"github.com/PubApiADN/pkg/logger"
)

const (
	convertCurrencyURL                   = "https://api.cambio.today/v1/quotes/%s/%s/json?quantity=%v&key=%v"
	loggerErrorSendingRequest            = "Error sending request to api.cambio.today [Class: CurrencyConvertClient][Method:GetCurrency]"
	loggerErrorGettingRequestInformation = "Error getting information from api.cambio.today [Class: CurrencyConvertClient][Method:GetCurrency]"
	apiName                              = "api.cambio.today"
)

type CurrencyConvertClient struct {
	RestClient config.CustomRestClient
}

func (currencyConvertClient *CurrencyConvertClient) GetCurrency(currency model.Currency) (model.CurrencyConversion, error) {

	var currencyConversion model.CurrencyConversion
	url := fmt.Sprintf(convertCurrencyURL, currency.Source, currency.Target, currency.Quantity, config.GetCurrencyApiKey())

	err := currencyConvertClient.RestClient.Get(url, apiName, &currencyConversion)

	if err != nil {

		logger.Error(loggerErrorSendingRequest, err)
		return model.CurrencyConversion{}, err
	}

	return currencyConversion, nil
}
