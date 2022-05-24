package container

import (
	"github.com/PubApiADN/cmd/api/app/domain/service/beer"
)

func getCreateBeerService() beer.CreateBeerService {
	return &beer.CreateBeer{
		BeerRepository: getCreateBeerRepository(),
	}
}

func getListBeerService() beer.ListBeerService {
	return &beer.ListBeer{
		BeerRepository: getCreateBeerRepository(),
	}
}

func getBeerService() beer.GetBeerService {
	return &beer.GetBeer{
		BeerRepository: getCreateBeerRepository(),
	}
}
func getConvertCurrencyService() beer.ConvertCurrencyService {
	return &beer.ConvertCurrency{
		ConvertCurrencyClient: getConvertCurrencyClient(),
	}
}

func getBeerBoxPriceService() beer.GetBeerBoxPriceService {
	return &beer.GetBeerBoxPrice{}
}
