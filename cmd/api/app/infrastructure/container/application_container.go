package container

import (
	"github.com/PubApiADN/cmd/api/app/application/beer"
)

func getCreateBeerApplication() beer.CreateBeerApplication {
	return &beer.CreateBeer{CreateBeerService: getCreateBeerService()}
}

func getListBeerApplication() beer.ListBeerApplication {
	return &beer.ListBeer{ListBeerService: getListBeerService()}
}

func getBeerApplication() beer.GetBeerApplication {
	return &beer.GetBeer{GetBeerService: getBeerService()}
}
func getBeerBoxPriceApplication() beer.GetBeerBoxPriceApplication {
	return &beer.GetBeerBoxPrice{
		GetBeerService:         getBeerService(),
		ConvertCurrencyService: getConvertCurrencyService(),
		GetBeerBoxPriceService: getBeerBoxPriceService(),
	}
}
