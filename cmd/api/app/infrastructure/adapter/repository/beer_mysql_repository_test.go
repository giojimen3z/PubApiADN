package repository_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/PubApiADN/cmd/api/app/domain/exception"
	"github.com/PubApiADN/cmd/api/app/domain/model"
	"github.com/PubApiADN/cmd/api/app/infrastructure/adapter/repository"
	"github.com/PubApiADN/cmd/api/test/builder"
)

const (
	insertQueryBeer = "INSERT INTO beer"
	beerListQuery   = "SELECT (.+) FROM beer"
)
const (
	beerId = 1
)

var _ = Describe("Repository", func() {
	Context("Beer Mysql Repository", func() {
		var (
			db                  *sql.DB
			dbMock              sqlmock.Sqlmock
			beerMysqlRepository repository.BeerMysqlRepository
		)
		BeforeEach(func() {
			db, dbMock, _ = sqlmock.New()
			beerMysqlRepository = repository.BeerMysqlRepository{
				WriteClient:          db,
				ReadConnectionClient: db,
			}

		})
		When("a new valid beer request is received  and save in dba", func() {
			It("should return nil error", func() {

				beer := builder.NewBeerDataBuilder().Build()

				dbMock.ExpectBegin()
				dbMock.ExpectExec(insertQueryBeer).WillReturnResult(sqlmock.NewResult(1, 1))
				dbMock.ExpectCommit()
				err := beerMysqlRepository.Save(beer)

				Expect(err).Should(BeNil())
			})
		})
		When("a new valid beer request and failed the transaction", func() {
			It("should return  error", func() {
				transactionErrorMessage := "an error happened when execute the transaction"
				beer := builder.NewBeerDataBuilder().Build()
				errorOnTransaction := exception.InternalServerError{ErrMessage: transactionErrorMessage}

				dbMock.ExpectBegin().WillReturnError(errorOnTransaction)

				err := beerMysqlRepository.Save(beer)

				Expect(err).Should(Not(BeNil()))
				Expect(errorOnTransaction).Should(Equal(err))
			})
		})

		When("a new valid beer request and failed insert into dba", func() {
			It("should return  error", func() {
				transactionErrorMessage := "an error happened when try insert into dba"
				beer := builder.NewBeerDataBuilder().Build()
				errorOnInsert := exception.InternalServerError{ErrMessage: transactionErrorMessage}
				dbMock.ExpectBegin()
				dbMock.ExpectExec(insertQueryBeer).WillReturnError(errorOnInsert)

				err := beerMysqlRepository.Save(beer)

				Expect(err).Should(Not(BeNil()))
				Expect(errorOnInsert).Should(Equal(err))
			})
		})
		When("a new valid request is received  and get all beers from dba", func() {
			It("should return beerList and nil error", func() {
				beer := builder.NewBeerDataBuilder().Build()
				beerListExpected := []model.Beer{beer}
				rows := sqlmock.NewRows([]string{"id", "name", "brewery", "country", "price", "currency"}).AddRow(beer.BeerId, beer.Name, beer.Brewery, beer.Country, beer.Price, beer.Currency)
				dbMock.ExpectQuery(beerListQuery).WillReturnRows(rows)

				beerList, err := beerMysqlRepository.ListBeer()

				Expect(err).Should(BeNil())
				Expect(beerListExpected).Should(Equal(beerList))
			})
		})
		When("a new invalid  request is received", func() {
			It("should return  error and empty list", func() {
				errorExpected := "some type of parameters is not correct"
				beerListExpected := []model.Beer{}
				dbMock.ExpectQuery(beerListQuery).WillReturnError(errors.New(errorExpected))

				beerList, err := beerMysqlRepository.ListBeer()

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				Expect(beerListExpected).Should(Equal(beerList))

			})
		})
		When("a new request is received and failed scanning the fields", func() {
			It("should return  error and empty list", func() {
				errorExpected := "sql: expected 5 destination arguments in Scan, not 6"
				beer := builder.NewBeerDataBuilder().Build()
				rows := sqlmock.NewRows([]string{"id", "name", "brewery", "country", "price"}).AddRow(beer.BeerId, beer.Name, beer.Brewery, beer.Country, beer.Price)
				dbMock.ExpectQuery(beerListQuery).WillReturnRows(rows)

				beerList, err := beerMysqlRepository.ListBeer()

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				Expect(beerList).Should(BeNil())

			})
		})
		When("a new valid request is received  and get beer from dba", func() {
			It("should return one beer and nil error", func() {

				beerExpected := builder.NewBeerDataBuilder().Build()
				rows := sqlmock.NewRows([]string{"id", "name", "brewery", "country", "price", "currency"}).AddRow(beerExpected.BeerId, beerExpected.Name, beerExpected.Brewery, beerExpected.Country, beerExpected.Price, beerExpected.Currency)
				args := []driver.Value{beerId}
				dbMock.ExpectQuery(beerListQuery).WithArgs(args...).WillReturnRows(rows)

				beer, err := beerMysqlRepository.GetBeerByID(beerId)

				Expect(err).Should(BeNil())
				Expect(beerExpected).Should(Equal(beer))
			})
		})
		When("a new invalid request is received  and failed getting beer from dba", func() {
			It("should return error", func() {
				errorExpected := errors.New("sql: expected 5 destination arguments in Scan, not 6")
				beerMock := builder.NewBeerDataBuilder().Build()
				beerExpected := model.Beer{}
				rows := sqlmock.NewRows([]string{"id", "name", "brewery", "country", "price"}).AddRow(beerMock.BeerId, beerMock.Name, beerMock.Brewery, beerMock.Country, beerMock.Price)
				args := []driver.Value{beerId}
				dbMock.ExpectQuery(beerListQuery).WithArgs(args...).WillReturnRows(rows)

				beer, err := beerMysqlRepository.GetBeerByID(beerId)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected.Error()).Should(Equal(err.Error()))
				Expect(beerExpected).Should(Equal(beer))

			})
		})
		When("a new invalid request is received  and failed scan beer from dba", func() {
			It("should return error", func() {
				errorExpected := errors.New("some type of parameters is not correct")
				beerExpected := model.Beer{}
				dbMock.ExpectQuery(beerListQuery).WillReturnError(errorExpected)

				beer, err := beerMysqlRepository.GetBeerByID(beerId)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected.Error()).Should(Equal(err.Error()))
				Expect(beerExpected).Should(Equal(beer))

			})
		})
	})
})
