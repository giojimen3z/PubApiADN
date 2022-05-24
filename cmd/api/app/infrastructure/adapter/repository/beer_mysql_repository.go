package repository

import (
	"database/sql"
	"fmt"

	"github.com/PubApiADN/cmd/api/app/domain/exception"
	"github.com/PubApiADN/cmd/api/app/domain/model"
	"github.com/PubApiADN/cmd/api/app/infrastructure/config"
	"github.com/PubApiADN/pkg/logger"
)

const (
	loggerErrorTransactionBeer = "an error occurred  in the  transaction to save the beer: %s   [Class: BeerMysqlRepository][Method:Save]"
	loggerErrorSavingBeer      = "an error occurred  in the  transaction to save the beer, Error:%s   [Class: BeerMysqlRepository][Method:Save]"
	loggerErrorSyntaxQuery     = "query syntax error [Class: BeerMysqlRepository][Method:%s]"
	loggerErrorScanningFields  = "error when scanning the information [Class: BeerMysqlRepository][Method:%s]"
	queryToSaveBeer            = `INSERT INTO beer (id, name, brewery, country, price, currency)VALUES(?,?,?,?,?,?)`
	queryGetAllBeer            = `SELECT * FROM beer`
	queryGetBeer               = `SELECT * FROM beer WHERE id = ?`
	listBeerMethodName         = "ListBeer"
	getBeerByIDMethodName      = "GetBeerByID"
)

// BeerMysqlRepository represent the mysql repository
type BeerMysqlRepository struct {
	WriteClient          *sql.DB
	ReadConnectionClient *sql.DB
}

func (beerMysqlRepository *BeerMysqlRepository) Save(beer model.Beer) (err error) {
	var tx *sql.Tx

	defer func() {
		config.CloseConnections(err, tx, nil, nil)
	}()
	tx, err = beerMysqlRepository.WriteClient.Begin()
	if err != nil {

		logger.Error(fmt.Sprintf(loggerErrorTransactionBeer, beer.Name), err)
		return exception.InternalServerError{ErrMessage: err.Error()}
	}
	_, err = beerMysqlRepository.WriteClient.Exec(queryToSaveBeer,
		beer.BeerId,
		beer.Name,
		beer.Brewery,
		beer.Country,
		beer.Price,
		beer.Currency)

	if err != nil {
		logger.Error(fmt.Sprintf(loggerErrorSavingBeer, err.Error()), err)
		return exception.InternalServerError{ErrMessage: err.Error()}
	}

	return err
}

func (beerMysqlRepository *BeerMysqlRepository) ListBeer() (beersList []model.Beer, err error) {
	var rowsBeer *sql.Rows

	rowsBeer, err = beerMysqlRepository.ReadConnectionClient.Query(queryGetAllBeer)
	if err != nil {
		logger.Error(fmt.Sprintf(loggerErrorSyntaxQuery, listBeerMethodName), err)
		return []model.Beer{}, err
	}

	defer rowsBeer.Close()

	for rowsBeer.Next() {
		beer := model.Beer{}
		err = rowsBeer.Scan(
			&beer.BeerId,
			&beer.Name,
			&beer.Brewery,
			&beer.Country,
			&beer.Price,
			&beer.Currency,
		)
		if err != nil {
			logger.Error(fmt.Sprintf(loggerErrorScanningFields, listBeerMethodName), err)
			return
		}
		beersList = append(beersList, beer)
	}
	err = rowsBeer.Err()

	return beersList, err
}

func (beerMysqlRepository *BeerMysqlRepository) GetBeerByID(id int64) (beer model.Beer, err error) {
	var rowsBeer *sql.Rows

	rowsBeer, err = beerMysqlRepository.ReadConnectionClient.Query(queryGetBeer, id)

	if err != nil {
		logger.Error(fmt.Sprintf(loggerErrorSyntaxQuery, getBeerByIDMethodName), err)
		return model.Beer{}, err
	}

	defer rowsBeer.Close()

	if rowsBeer.Next() {
		err = rowsBeer.Scan(&beer.BeerId,
			&beer.Name,
			&beer.Brewery,
			&beer.Country,
			&beer.Price,
			&beer.Currency,
		)
		if err != nil {
			logger.Error(fmt.Sprintf(loggerErrorScanningFields, getBeerByIDMethodName), err)
			return model.Beer{}, err
		}
	}

	return beer, nil
}
