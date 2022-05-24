package beer

import (
	"github.com/PubApiADN/cmd/api/app/domain/model"
	"github.com/PubApiADN/cmd/api/app/domain/port"
	"github.com/PubApiADN/pkg/apierrors"
	"github.com/PubApiADN/pkg/logger"
)

const (
	errorConvertCurrency       = "error converting the currency into repository"
	loggerErrorConvertCurrency = "error converting the currency into repository [Class: ConvertCurrencyService][Method:ConvertCurrency]"
)

type ConvertCurrencyService interface {
	// ConvertCurrency Send to repository the currency
	ConvertCurrency(currency model.Currency) (model.CurrencyConversion, apierrors.ApiError)
}

type ConvertCurrency struct {
	ConvertCurrencyClient port.CurrencyClient
}

func (convertCurrency *ConvertCurrency) ConvertCurrency(currency model.Currency) (model.CurrencyConversion, apierrors.ApiError) {

	currencyConversion, errorRepository := convertCurrency.ConvertCurrencyClient.GetCurrency(currency)

	if errorRepository != nil {
		logger.Error(loggerErrorConvertCurrency, errorRepository)
		err := apierrors.NewNotFoundApiError(errorConvertCurrency)
		return model.CurrencyConversion{}, err
	}

	return currencyConversion, nil

}
