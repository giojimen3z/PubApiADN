package client_test

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/PubApiADN/cmd/api/app/domain/model"
	"github.com/PubApiADN/cmd/api/app/infrastructure/adapter/client"
	"github.com/PubApiADN/cmd/api/app/infrastructure/config"
	"github.com/PubApiADN/cmd/api/test/builder"
)

const (
	beerId = 1
)

var _ = Describe("Client", func() {
	Context("Rest client", func() {
		var (
			restyClient           = resty.New()
			currencyConvertClient client.CurrencyConvertClient
		)

		var _ = BeforeSuite(func() {

			httpmock.ActivateNonDefault(restyClient.GetClient())
		})

		BeforeEach(func() {
			httpmock.Reset()
			_ = os.Setenv("SCOPE", "local")

			currencyConvertClient = client.CurrencyConvertClient{
				RestClient: config.CustomRestClient{},
			}

		})

		AfterEach(func() {
			os.Clearenv()
		})

		var _ = AfterSuite(func() {
			httpmock.DeactivateAndReset()
		})

		When("a new valid  request is received  and convert currency successfully", func() {
			It("should return CurrencyConversion struct and nil error", func() {
				fakeUrl := "https://api.cambio.today/v1/quotes/COP/USD/json?quantity=1500&key=6392|h_2OeBxS2ibfZ^D1cA1o_3cYBQNUD*Pm"
				currency := builder.NewCurrencyDataBuilder().Build()
				currencyConversionExpected := builder.NewCurrencyConversionDataBuilder().Build()
				fixture, _ := json.Marshal(currencyConversionExpected)
				responder := httpmock.NewStringResponder(200, string(fixture))
				httpmock.RegisterResponder("GET", fakeUrl, responder)
				restyClient.R().SetResult(currencyConversionExpected).Get(fakeUrl)

				currencyConversion, err := currencyConvertClient.GetCurrency(currency)

				Expect(err).Should(BeNil())
				Expect(currencyConversionExpected).Should(Equal(currencyConversion))

			})
		})
		When("a new  request is received  and failed convert currency", func() {
			It("should return CurrencyConversion struct empty and  error", func() {
				fakeUrl := "https://api.cambio.today"
				currency := model.Currency{}
				currencyConversionExpected := model.CurrencyConversion{}
				errorExpected := errors.New("the body of the petition is incorrect, check and try again")
				responder := httpmock.NewErrorResponder(errorExpected)
				httpmock.RegisterResponder("GET", fakeUrl, responder)
				restyClient.R().SetError(errorExpected).SetResult(currencyConversionExpected).Get(fakeUrl)

				currencyConversion, err := currencyConvertClient.GetCurrency(currency)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected.Error()).Should(Equal(err.Error()))
				Expect(currencyConversionExpected).Should(Equal(currencyConversion))
			})
		})

	})
})
