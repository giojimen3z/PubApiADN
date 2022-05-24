package builder

import "github.com/PubApiADN/cmd/api/app/domain/model"

type CurrencyDataBuilder struct {
	updated  string
	source   string
	target   string
	value    float64
	quantity float64
	amount   float64
}

func NewCurrencyDataBuilder() *CurrencyDataBuilder {
	return &CurrencyDataBuilder{
		updated:  "2020-11-28T00:11:20",
		source:   "COP",
		target:   "USD",
		value:    0.00028,
		quantity: 1500.0,
		amount:   0.4161,
	}
}

func (builder *CurrencyDataBuilder) Build() model.Currency {
	return model.Currency{
		Updated:  builder.updated,
		Source:   builder.source,
		Target:   builder.target,
		Value:    builder.value,
		Quantity: builder.quantity,
		Amount:   builder.amount,
	}
}
