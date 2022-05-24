package beer_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	mockParameter "github.com/stretchr/testify/mock"

	"github.com/PubApiADN/cmd/api/app/application/beer"
	"github.com/PubApiADN/cmd/api/app/domain/model"
	beer2 "github.com/PubApiADN/cmd/api/app/domain/service/beer"
	"github.com/PubApiADN/cmd/api/test/builder"
	"github.com/PubApiADN/cmd/api/test/mock"
)

const (
	currency = "USD"
	quantity = 1
)

var _ = Describe("Handler", func() {
	Context("Get Beer Box Price", func() {
		var (
			repositoryMock         *mock.BeerRepositoryMock
			clientMock             *mock.CurrencyClientMock
			getBeerBoxPriceUseCase beer.GetBeerBoxPrice
		)
		BeforeEach(func() {
			repositoryMock = new(mock.BeerRepositoryMock)
			clientMock = new(mock.CurrencyClientMock)
			getBeerService := &beer2.GetBeer{
				BeerRepository: repositoryMock,
			}
			convertCurrencyService := &beer2.ConvertCurrency{
				ConvertCurrencyClient: clientMock,
			}
			getBeerBoxPrice := &beer2.GetBeerBoxPrice{}
			getBeerBoxPriceUseCase = beer.GetBeerBoxPrice{
				GetBeerService:         getBeerService,
				ConvertCurrencyService: convertCurrencyService,
				GetBeerBoxPriceService: getBeerBoxPrice,
			}
		})
		When("a new valid  request is received", func() {
			It("should return beer box price  and nil error", func() {
				currencyConversion := builder.NewCurrencyConversionDataBuilder().Build()
				beerExpected := builder.NewBeerDataBuilder().Build()
				beerBoxExpected := builder.NewBeerBoxDataBuilder().Build()
				repositoryMock.On("GetBeerByID", mockParameter.Anything).Return(beerExpected, nil)
				clientMock.On("GetCurrency", mockParameter.Anything).Return(currencyConversion, nil)

				beerBox, err := getBeerBoxPriceUseCase.Handler(beerId, currency, quantity)

				Expect(err).Should(BeNil())
				Expect(beerBoxExpected).Should(Equal(beerBox))
				repositoryMock.AssertExpectations(GinkgoT())
				clientMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid  request is received and getBeerService Failed", func() {
			It("should return error", func() {
				errorRepository := errors.New("some type of parameters isn't correct")
				beerExpected := model.Beer{}
				beerBoxExpected := model.BeerBox{}
				errorExpected := "Message: The beer id:1 isn´t exists;Error Code: not_found;Status: 404;Cause: []"
				repositoryMock.On("GetBeerByID", mockParameter.Anything).Return(beerExpected, errorRepository)

				beerBox, err := getBeerBoxPriceUseCase.Handler(beerId, currency, quantity)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				Expect(beerBoxExpected).Should(Equal(beerBox))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid  request is received and convertCurrencyService Failed", func() {
			It("should return error", func() {
				currencyConversion := model.CurrencyConversion{}
				errorRepository := errors.New("error converting the currency into repository")
				beerExpected := builder.NewBeerDataBuilder().Build()
				beerBoxExpected := model.BeerBox{}
				errorExpected := "Message: error converting the currency into repository;Error Code: not_found;Status: 404;Cause: []"
				repositoryMock.On("GetBeerByID", mockParameter.Anything).Return(beerExpected, nil)
				clientMock.On("GetCurrency", mockParameter.Anything).Return(currencyConversion, errorRepository)

				beerBox, err := getBeerBoxPriceUseCase.Handler(beerId, currency, quantity)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				Expect(beerBoxExpected).Should(Equal(beerBox))
				repositoryMock.AssertExpectations(GinkgoT())
				clientMock.AssertExpectations(GinkgoT())
			})
		})
	})
})
