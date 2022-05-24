package app

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/PubApiADN/cmd/api/app/infrastructure/config"
	"github.com/PubApiADN/cmd/api/app/infrastructure/container"
	controller "github.com/PubApiADN/cmd/api/app/infrastructure/controller/health"
)

func MapUrls(router *gin.Engine) {
	prefixScope := config.GetRoutePrefix()
	router.GET("/ping", controller.PingController.Ping)
	prefix := fmt.Sprintf("%s/PubApi/", prefixScope)

	baseUrl := router.Group(prefix)
	beer := baseUrl.Group("/Beer")
	beers := baseUrl.Group("/Beers")

	beer.POST("", container.GetCreateBeerController().MakeCreateBeer)
	beer.GET(":id", container.GetBeerController().MakeGetBeer)
	beer.GET(":id/BoxPrice", container.GetBeerBoxPriceController().MakeGetBeerBoxPrice)
	beers.GET("", container.GetListBeerController().MakeListBeer)

}
