package port

import "github.com/PubApiADN/cmd/api/app/domain/model"

// BeerRepository  use for all transactions about beer
type BeerRepository interface {
	//Save persist the beer data
	Save(beer model.Beer) (err error)
	//ListBeer get all beers from persistence
	ListBeer() (beersList []model.Beer, err error)
	//GetBeerByID get  beers for id from persistence
	GetBeerByID(id int64) (beer model.Beer, err error)
}
