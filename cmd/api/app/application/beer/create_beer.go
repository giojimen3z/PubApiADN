package beer

import (
	"github.com/PubApiADN/cmd/api/app/domain/model"
	"github.com/PubApiADN/cmd/api/app/domain/service/beer"
	"github.com/PubApiADN/pkg/apierrors"
)

// CreateBeerApplication is the initial flow entry to create one beer
type CreateBeerApplication interface {
	// Handler is the input for access to create one beer
	Handler(beer model.Beer) apierrors.ApiError
}
type CreateBeer struct {
	CreateBeerService beer.CreateBeerService
}

func (createBeer *CreateBeer) Handler(beer model.Beer) apierrors.ApiError {
	return createBeer.CreateBeerService.CreateBeer(beer)
}
