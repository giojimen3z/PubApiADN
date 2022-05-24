package beer

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/PubApiADN/cmd/api/app/application/beer"
)

// ListCreateBeerController  used for inject the use case
type ListBeerController struct {
	ListBeerApplication beer.ListBeerApplication
}

func (listBeerController *ListBeerController) MakeListBeer(context *gin.Context) {

	BeerList, err := listBeerController.ListBeerApplication.Handler()

	if err != nil {
		context.JSON(err.Status(), err)
	}

	context.JSON(http.StatusOK, BeerList)

}
