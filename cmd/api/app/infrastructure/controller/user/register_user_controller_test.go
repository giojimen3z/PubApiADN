package user_test

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

	userApplication "github.com/PubApiADN/cmd/api/app/application/user"
	"github.com/PubApiADN/cmd/api/app/domain/model"
	userService "github.com/PubApiADN/cmd/api/app/domain/service/user"
	userController "github.com/PubApiADN/cmd/api/app/infrastructure/controller/user"
	"github.com/PubApiADN/cmd/api/test/builder"
	"github.com/PubApiADN/cmd/api/test/mock"
)

var _ = Describe("User Controller", func() {
	Context("Register User", func() {
		var (
			user                   model.User
			registerUserController userController.RegisterUserController
			context                *gin.Context
			repositoryMock         *mock.UserRepositoryMock
			recorder               *httptest.ResponseRecorder
		)
		BeforeEach(func() {
			_ = os.Setenv("SCOPE", "local")
			user = builder.NewUserDataBuilder().Build()
			recorder = httptest.NewRecorder()
			context, _ = gin.CreateTestContext(recorder)
			repositoryMock = new(mock.UserRepositoryMock)
			registerUserService := &userService.RegisterUser{
				UserRepository: repositoryMock,
			}
			registerUserUseCase := &userApplication.RegisterUser{
				RegisterUserService: registerUserService,
			}
			registerUserController = userController.RegisterUserController{
				RegisterUserApplication: registerUserUseCase,
			}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid request is received", func() {
			It("should return 201 code", func() {
				body, _ := json.Marshal(user)
				context.Request, _ = http.NewRequest("POST", "/testing", strings.NewReader(string(body)))
				repositoryMock.On("Save", user).Return(nil)
				expectMessage := "\"the user Gio was created successfully\""

				registerUserController.MakeRegisterUser(context)

				Expect(http.StatusCreated).To(Equal(recorder.Code))
				Expect(expectMessage).Should(Equal(recorder.Body.String()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid request is received", func() {
			It("should return 400 code", func() {

				context.Request, _ = http.NewRequest("POST", "/testing", strings.NewReader(string("")))
				errorExpected := "{\"message\":\"invalid request\",\"error\":\"bad_request\",\"status\":400,\"cause\":[]}"

				registerUserController.MakeRegisterUser(context)

				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
				Expect(errorExpected).Should(Equal(recorder.Body.String()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})

		When("a new valid request is received but with  invalid id", func() {
			It("should return 409 code", func() {
				body, _ := json.Marshal(user)
				context.Request, _ = http.NewRequest("POST", "/testing", strings.NewReader(string(body)))
				errorMock := errors.New("the id:1 is invalid")
				repositoryMock.On("Save", user).Return(errorMock)
				errorExpected := "{\"message\":\"User id:1 already exists\",\"error\":\"Conflict\",\"status\":409,\"cause\":null}"

				registerUserController.MakeRegisterUser(context)

				Expect(http.StatusConflict).To(Equal(recorder.Code))
				Expect(errorExpected).Should(Equal(recorder.Body.String()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})

	})
})
