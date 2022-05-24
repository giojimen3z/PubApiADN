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
	beerId = 1
)

var _ = Describe("Handler", func() {
	Context("Get Beer", func() {
		var (
			repositoryMock *mock.BeerRepositoryMock
			getBeerUseCase beer.GetBeer
		)
		BeforeEach(func() {
			repositoryMock = new(mock.BeerRepositoryMock)
			getBeerService := &beer2.GetBeer{
				BeerRepository: repositoryMock,
			}
			getBeerUseCase = beer.GetBeer{
				GetBeerService: getBeerService,
			}
		})
		When("a new valid  request is received", func() {
			It("should return beer  and nil error", func() {
				beerExpected := builder.NewBeerDataBuilder().Build()
				repositoryMock.On("GetBeerByID", mockParameter.Anything).Return(beerExpected, nil)

				beer, err := getBeerUseCase.Handler(beerId)

				Expect(err).Should(BeNil())
				Expect(beerExpected).Should(Equal(beer))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid  request is received", func() {
			It("should return error", func() {
				errorRepository := errors.New("some type of parameters isn't correct")
				beerExpected := model.Beer{}
				errorExpected := "Message: The beer id:1 isn´t exists;Error Code: not_found;Status: 404;Cause: []"
				repositoryMock.On("GetBeerByID", mockParameter.Anything).Return(beerExpected, errorRepository)

				beer, err := getBeerUseCase.Handler(beerId)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				Expect(beerExpected).Should(Equal(beer))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
	})
})
