package builder

import "github.com/PubApiADN/cmd/api/app/domain/model"

type CurrencyConversionDataBuilder struct {
	currency model.Currency
	status   string
}

func NewCurrencyConversionDataBuilder() *CurrencyConversionDataBuilder {
	return &CurrencyConversionDataBuilder{
		currency: NewCurrencyDataBuilder().Build(),
		status:   "OK",
	}
}

func (builder *CurrencyConversionDataBuilder) Build() model.CurrencyConversion {
	return model.CurrencyConversion{
		Currency: builder.currency,
		Status:   builder.status,
	}
}
