package beer

import (
	"github.com/PubApiADN/cmd/api/app/domain/model"
)

const (
	Zero            = 0
	defaultQuantity = 6
)

type GetBeerBoxPriceService interface {
	// ConvertCurrency Send to repository the currency
	GetBeerBoxPrice(quantity int64, currencyConversion model.CurrencyConversion) model.BeerBox
}

type GetBeerBoxPrice struct{}

func (getBeerBoxPrice *GetBeerBoxPrice) GetBeerBoxPrice(quantity int64, currencyConversion model.CurrencyConversion) model.BeerBox {

	boxQuantity := getBeerBoxPrice.getQuantityBox(quantity)

	price := float64(boxQuantity) * currencyConversion.Currency.Amount

	beerBox := model.BeerBox{
		Price: price,
	}

	return beerBox
}

func (getBeerBoxPrice *GetBeerBoxPrice) getQuantityBox(quantity int64) int64 {
	if quantity == Zero {
		quantity = defaultQuantity
	} else {

		quantity = quantity * 6
	}
	return quantity
}
