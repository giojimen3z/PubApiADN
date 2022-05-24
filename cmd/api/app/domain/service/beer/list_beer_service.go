package beer

import (
	"errors"

	"github.com/PubApiADN/cmd/api/app/domain/model"
	"github.com/PubApiADN/cmd/api/app/domain/port"
	"github.com/PubApiADN/pkg/apierrors"
	"github.com/PubApiADN/pkg/logger"
)

const (
	ZeroFields            = 0
	errorListBeer         = "Error getting the beers from repository"
	logErrorListBeer      = "Error getting the beers from repository  [Class: ListBeerService][Method:ListBeer]"
	errorEmptyListBeer    = "Error with the  information received, the Beers are empty"
	logEmptyErrorListBeer = "Error with the  information received, the Beers are empty  [Class: ListBeerService][Method:ListBeer]"
)

type ListBeerService interface {
	// listBeer Get all beers from repository
	ListBeer() ([]model.Beer, apierrors.ApiError)
}

type ListBeer struct {
	BeerRepository port.BeerRepository
}

func (listBeer *ListBeer) ListBeer() ([]model.Beer, apierrors.ApiError) {

	beers, errorRepository := listBeer.BeerRepository.ListBeer()

	if errorRepository != nil {
		logger.Error(logErrorListBeer, errorRepository)
		err := apierrors.NewBadRequestApiError(errorListBeer)
		return []model.Beer{}, err
	}

	if len(beers) == ZeroFields {
		errorEmptyBeers := errors.New(errorEmptyListBeer)
		logger.Error(logErrorListBeer, errorRepository)
		err := apierrors.NewInternalServerApiError(errorEmptyBeers.Error(), errorEmptyBeers)
		return []model.Beer{}, err
	}

	return beers, nil

}
