package beer_test

import (
	"errors"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	mockParameter "github.com/stretchr/testify/mock"

	"github.com/PubApiADN/cmd/api/app/domain/model"
	"github.com/PubApiADN/cmd/api/app/domain/service/beer"
	"github.com/PubApiADN/cmd/api/test/builder"
	"github.com/PubApiADN/cmd/api/test/mock"
)

const (
	beerId = 1
	zeroId = 0
)

var _ = Describe("Service", func() {
	Context("Get Beer", func() {
		var (
			repositoryMock *mock.BeerRepositoryMock
			getBeerService beer.GetBeer
		)
		BeforeEach(func() {
			repositoryMock = new(mock.BeerRepositoryMock)
			getBeerService = beer.GetBeer{
				BeerRepository: repositoryMock,
			}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid request is received", func() {
			It("should return one beer and nil error", func() {

				beerExpected := builder.NewBeerDataBuilder().Build()
				repositoryMock.On("GetBeerByID", mockParameter.Anything).Return(beerExpected, nil)

				beer, err := getBeerService.GetBeer(beerId)

				Expect(err).Should(BeNil())
				Expect(beerExpected).Should(Equal(beer))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid request is received", func() {
			It("should return error", func() {
				errorRepository := errors.New("some type of parameters isn't correct")
				beerExpected := model.Beer{}
				errorExpected := "Message: The beer id:1 isn´t exists;Error Code: not_found;Status: 404;Cause: []"
				repositoryMock.On("GetBeerByID", mockParameter.Anything).Return(beerExpected, errorRepository)

				beer, err := getBeerService.GetBeer(beerId)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				Expect(beerExpected).Should(Equal(beer))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid request is received", func() {
			It("should return error", func() {
				beerExpected := builder.NewBeerDataBuilderWithZeroID().Build()
				errorExpected := "Message: The beer id:0 isn´t exists;Error Code: not_found;Status: 404;Cause: []"
				repositoryMock.On("GetBeerByID", mockParameter.Anything).Return(beerExpected, nil)

				beer, err := getBeerService.GetBeer(zeroId)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				Expect(beerExpected).Should(Equal(beer))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
	})
})
