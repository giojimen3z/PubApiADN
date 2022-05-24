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

var _ = Describe("Service", func() {
	Context("Convert Currency", func() {
		var (
			clientMock                *mock.CurrencyClientMock
			getConvertCurrencyService beer.ConvertCurrency
		)
		BeforeEach(func() {
			clientMock = new(mock.CurrencyClientMock)
			getConvertCurrencyService = beer.ConvertCurrency{
				ConvertCurrencyClient: clientMock,
			}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid request is received", func() {
			It("should return currency conversion and nil error", func() {

				currency := builder.NewCurrencyDataBuilder().Build()
				currencyConversionExpected := builder.NewCurrencyConversionDataBuilder().Build()
				clientMock.On("GetCurrency", mockParameter.Anything).Return(currencyConversionExpected, nil)

				currencyConversion, err := getConvertCurrencyService.ConvertCurrency(currency)

				Expect(err).Should(BeNil())
				Expect(currencyConversionExpected).Should(Equal(currencyConversion))
				clientMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new valid request is received", func() {
			It("should return  error", func() {
				errorRepository := errors.New("error converting the currency into repository")
				errorExpected := "Message: error converting the currency into repository;Error Code: not_found;Status: 404;Cause: []"
				currency := builder.NewCurrencyDataBuilder().Build()
				currencyConversionExpected := model.CurrencyConversion{}
				clientMock.On("GetCurrency", mockParameter.Anything).Return(currencyConversionExpected, errorRepository)

				currencyConversion, err := getConvertCurrencyService.ConvertCurrency(currency)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				Expect(currencyConversionExpected).Should(Equal(currencyConversion))
				clientMock.AssertExpectations(GinkgoT())
			})
		})
	})
})
