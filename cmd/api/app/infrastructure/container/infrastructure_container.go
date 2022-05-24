package container

import (
	"database/sql"

	"github.com/PubApiADN/cmd/api/app/domain/port"
	"github.com/PubApiADN/cmd/api/app/infrastructure/adapter/client"
	"github.com/PubApiADN/cmd/api/app/infrastructure/adapter/repository"
	"github.com/PubApiADN/cmd/api/app/infrastructure/config"
	"github.com/PubApiADN/cmd/api/app/infrastructure/controller/beer"
)

func GetCreateBeerController() *beer.CreateBeerController {
	return &beer.CreateBeerController{CreateBeerApplication: getCreateBeerApplication()}
}

func GetListBeerController() *beer.ListBeerController {
	return &beer.ListBeerController{ListBeerApplication: getListBeerApplication()}
}
func GetBeerController() *beer.GetBeerController {
	return &beer.GetBeerController{GetBeerApplication: getBeerApplication()}
}
func GetBeerBoxPriceController() *beer.GetBeerBoxPriceController {
	return &beer.GetBeerBoxPriceController{GetBeerBoxPriceApplication: getBeerBoxPriceApplication()}
}

func getCreateBeerRepository() port.BeerRepository {
	return &repository.BeerMysqlRepository{
		WriteClient:          getWriteConnectionClient(),
		ReadConnectionClient: getReadConnectionClient(),
	}
}

func getConvertCurrencyClient() port.CurrencyClient {
	return &client.CurrencyConvertClient{
		RestClient: getRestConnectionClient(),
	}
}
func getWriteConnectionClient() *sql.DB {
	conn, _ := config.GetWriteConnection()
	return conn
}

func getReadConnectionClient() *sql.DB {
	conn, _ := config.GetReadConnection()
	return conn
}

func getRestConnectionClient() config.CustomRestClient {

	return config.CustomRestClient{
	}
}
