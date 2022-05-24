package beer

import (
	"fmt"

	"github.com/PubApiADN/cmd/api/app/domain/model"
	"github.com/PubApiADN/cmd/api/app/domain/port"
	"github.com/PubApiADN/pkg/apierrors"
	"github.com/PubApiADN/pkg/logger"
)

const (
	errorIDNotExist    = "The beer id:%v isn´t exists"
	logErrorIDNotExist = "The beer id:%v isn´t exists [Class: GetBeerService][Method:GetBeer]"
	ZeroID             = 0
)

type GetBeerService interface {
	// GetBeer Send to repository the id
	GetBeer(id int64) (model.Beer, apierrors.ApiError)
}

type GetBeer struct {
	BeerRepository port.BeerRepository
}

func (getBeer *GetBeer) GetBeer(id int64) (model.Beer, apierrors.ApiError) {

	beer, errorRepository := getBeer.BeerRepository.GetBeerByID(id)

	if errorRepository != nil {
		logger.Error(fmt.Sprintf(logErrorIDNotExist, id), errorRepository)
		err := apierrors.NewNotFoundApiError(fmt.Sprintf(errorIDNotExist, id))
		return model.Beer{}, err
	}

	if beer.BeerId == ZeroID {
		logger.Error(fmt.Sprintf(logErrorIDNotExist, id), errorRepository)
		err := apierrors.NewNotFoundApiError(fmt.Sprintf(errorIDNotExist, id))
		return model.Beer{}, err
	}

	return beer, nil
}
