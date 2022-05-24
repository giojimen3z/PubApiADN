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

	beer2 "github.com/PubApiADN/cmd/api/app/application/beer"
	"github.com/PubApiADN/cmd/api/app/domain/model"
	beer3 "github.com/PubApiADN/cmd/api/app/domain/service/beer"
	beerController "github.com/PubApiADN/cmd/api/app/infrastructure/controller/beer"
	"github.com/PubApiADN/cmd/api/test/builder"
	"github.com/PubApiADN/cmd/api/test/mock"
)

var _ = Describe("Beer Controller", func() {
	Context("List Beers", func() {
		var (
			beer               model.Beer
			listBeerController beerController.ListBeerController
			context            *gin.Context
			repositoryMock     *mock.BeerRepositoryMock
			recorder           *httptest.ResponseRecorder
		)
		BeforeEach(func() {
			_ = os.Setenv("SCOPE", "local")
			beer = builder.NewBeerDataBuilder().Build()
			recorder = httptest.NewRecorder()
			context, _ = gin.CreateTestContext(recorder)
			repositoryMock = new(mock.BeerRepositoryMock)
			listBeerService := &beer3.ListBeer{
				BeerRepository: repositoryMock,
			}
			listBeerUseCase := &beer2.ListBeer{
				ListBeerService: listBeerService,
			}
			listBeerController = beerController.ListBeerController{
				ListBeerApplication: listBeerUseCase,
			}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid request is received", func() {
			It("should return 200 code and list of beers", func() {

				beerList := []model.Beer{beer}
				bodyExpected, _ := json.Marshal(beerList)
				context.Request, _ = http.NewRequest("GET", "/testing", strings.NewReader(string("")))
				repositoryMock.On("ListBeer").Return(beerList, nil)

				listBeerController.MakeListBeer(context)

				Expect(http.StatusOK).To(Equal(recorder.Code))
				Expect(string(bodyExpected)).Should(Equal(recorder.Body.String()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new valid request is received", func() {
			It("should return 400 code", func() {

				beerList := []model.Beer{}
				errorRepository := errors.New("some type of parameters isn't correct")
				errorExpected := "{\"message\":\"Error getting the beers from repository\",\"error\":\"bad_request\",\"status\":400,\"cause\":[]}[]"
				context.Request, _ = http.NewRequest("GET", "/testing", strings.NewReader(string("")))
				repositoryMock.On("ListBeer").Return(beerList, errorRepository)

				listBeerController.MakeListBeer(context)

				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
				Expect(errorExpected).Should(Equal(recorder.Body.String()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})

		When("a new valid request is received", func() {
			It("should return 500 code", func() {

				beerList := []model.Beer{}
				errorExpected := "{\"message\":\"Error with the  information received, the Beers are empty\",\"error\":\"internal_server_error\",\"status\":500,\"cause\":[\"Error with the  information received, the Beers are empty\"]}[]"
				context.Request, _ = http.NewRequest("GET", "/testing", strings.NewReader(string("")))
				repositoryMock.On("ListBeer").Return(beerList, nil)

				listBeerController.MakeListBeer(context)

				Expect(http.StatusInternalServerError).To(Equal(recorder.Code))
				Expect(errorExpected).Should(Equal(recorder.Body.String()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
	})
})
