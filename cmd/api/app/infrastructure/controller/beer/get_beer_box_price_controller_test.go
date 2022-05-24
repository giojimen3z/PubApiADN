package beer_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	mockParameter "github.com/stretchr/testify/mock"

	beer2 "github.com/PubApiADN/cmd/api/app/application/beer"
	"github.com/PubApiADN/cmd/api/app/domain/model"
	beer3 "github.com/PubApiADN/cmd/api/app/domain/service/beer"
	beerController "github.com/PubApiADN/cmd/api/app/infrastructure/controller/beer"
	"github.com/PubApiADN/cmd/api/test/builder"
	"github.com/PubApiADN/cmd/api/test/mock"
)

var _ = Describe("Beer Controller", func() {
	Context("Get Beer Box Price", func() {
		var (
			beer                      model.Beer
			beerBox                   model.BeerBox
			currencyConversion        model.CurrencyConversion
			getBeerBoxPriceController beerController.GetBeerBoxPriceController
			context                   *gin.Context
			repositoryMock            *mock.BeerRepositoryMock
			clientMock                *mock.CurrencyClientMock
			recorder                  *httptest.ResponseRecorder
		)
		BeforeEach(func() {
			_ = os.Setenv("SCOPE", "local")
			beer = builder.NewBeerDataBuilder().Build()
			beerBox = builder.NewBeerBoxDataBuilder().WithPrice(2.4966).Build()
			currencyConversion = builder.NewCurrencyConversionDataBuilder().Build()
			recorder = httptest.NewRecorder()
			context, _ = gin.CreateTestContext(recorder)
			repositoryMock = new(mock.BeerRepositoryMock)
			clientMock = new(mock.CurrencyClientMock)
			getBeerService := &beer3.GetBeer{
				BeerRepository: repositoryMock,
			}
			convertCurrencyService := &beer3.ConvertCurrency{
				ConvertCurrencyClient: clientMock,
			}
			getBeerBoxPrice := &beer3.GetBeerBoxPrice{}
			getBeerBoxPriceUseCase := &beer2.GetBeerBoxPrice{
				GetBeerService:         getBeerService,
				ConvertCurrencyService: convertCurrencyService,
				GetBeerBoxPriceService: getBeerBoxPrice,
			}
			getBeerBoxPriceController = beerController.GetBeerBoxPriceController{
				GetBeerBoxPriceApplication: getBeerBoxPriceUseCase,
			}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid request is received", func() {
			It("should return 200 code and  beer box price", func() {

				bodyExpected, _ := json.Marshal(beerBox)
				context.Request, _ = http.NewRequest("GET", "/testing", strings.NewReader(string("")))
				repositoryMock.On("GetBeerByID", mockParameter.Anything).Return(beer, nil)
				clientMock.On("GetCurrency", mockParameter.Anything).Return(currencyConversion, nil)

				getBeerBoxPriceController.MakeGetBeerBoxPrice(context)

				Expect(http.StatusOK).To(Equal(recorder.Code))
				Expect(string(bodyExpected)).Should(Equal(recorder.Body.String()))
				repositoryMock.AssertExpectations(GinkgoT())
				clientMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid request is received", func() {
			It("should return 404 code and one error", func() {

				ErrorExpected := "{\"message\":\"The beer id:0 isn´t exists\",\"error\":\"not_found\",\"status\":404,\"cause\":[]}"
				errorRepository := errors.New("some type of parameters isn't correct")
				context.Request, _ = http.NewRequest("GET", "/testing", strings.NewReader(string("")))
				repositoryMock.On("GetBeerByID", mockParameter.Anything).Return(beer, errorRepository)

				getBeerBoxPriceController.MakeGetBeerBoxPrice(context)

				Expect(http.StatusNotFound).To(Equal(recorder.Code))
				Expect(ErrorExpected).Should(Equal(recorder.Body.String()))
				repositoryMock.AssertExpectations(GinkgoT())

			})
		})
	})
})
