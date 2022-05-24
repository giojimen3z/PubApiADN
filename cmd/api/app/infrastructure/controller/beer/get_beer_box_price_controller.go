package beer

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/PubApiADN/cmd/api/app/application/beer"
)

// GetBeerBoxPriceController  used for inject the use case
type GetBeerBoxPriceController struct {
	GetBeerBoxPriceApplication beer.GetBeerBoxPriceApplication
}

func (getBeerBoxPriceController *GetBeerBoxPriceController) MakeGetBeerBoxPrice(context *gin.Context) {

	beerID, currency, quantity := getBeerBoxPriceController.mapRequestedValues(context)

	beerBox, err := getBeerBoxPriceController.GetBeerBoxPriceApplication.Handler(beerID, currency, quantity)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, beerBox)
}

func (getBeerBoxPriceController *GetBeerBoxPriceController) mapRequestedValues(context *gin.Context) (int64, string, int64) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 16)
	beerID := int64(id)
	currency := strings.TrimSpace(context.Query("currency"))
	quantityRequested, _ := strconv.ParseUint(context.Query("quantity"), 10, 16)
	quantity := int64(quantityRequested)

	return beerID, currency, quantity
}
