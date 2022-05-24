package port

import "github.com/PubApiADN/cmd/api/app/domain/model"

//CurrencyClient  use for all transactions about currency conversion
type CurrencyClient interface {
	//GetCurrency convert the actual currency  for the currency requested
	GetCurrency(currency model.Currency) (model.CurrencyConversion, error)
}
