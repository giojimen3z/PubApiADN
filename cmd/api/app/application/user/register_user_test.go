package user_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	userApplication "github.com/PubApiADN/cmd/api/app/application/user"
	userService "github.com/PubApiADN/cmd/api/app/domain/service/user"
	"github.com/PubApiADN/cmd/api/test/builder"
	"github.com/PubApiADN/cmd/api/test/mock"
)

var _ = Describe("Handler", func() {
	Context("Register User", func() {
		var (
			repositoryMock      *mock.UserRepositoryMock
			registerUserUseCase userApplication.RegisterUser
		)
		BeforeEach(func() {
			repositoryMock = new(mock.UserRepositoryMock)
			userRegisterService := &userService.RegisterUser{
				UserRepository: repositoryMock,
			}
			registerUserUseCase = userApplication.RegisterUser{
				RegisterUserService: userRegisterService,
			}

		})

		When("a new valid user request is received", func() {
			It("should return nil error", func() {

				user := builder.NewUserDataBuilder().Build()
				repositoryMock.On("Save", user).Return(nil)

				err := registerUserUseCase.Handler(user)

				Expect(err).Should(BeNil())
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
		When("a new invalid user request is received", func() {
			It("should return error", func() {

				errorMock := errors.New("Error 1062: Duplicate entry '1' for key 'user.PRIMARY'")
				user := builder.NewUserDataBuilder().Build()
				repositoryMock.On("Save", user).Return(errorMock)
				errorExpected := "Message: User id:1 already exists;Error Code: Conflict;Status: 409;Cause: []"

				err := registerUserUseCase.Handler(user)

				Expect(err).Should(Not(BeNil()))
				Expect(errorExpected).Should(Equal(err.Error()))
				repositoryMock.AssertExpectations(GinkgoT())
			})
		})
	})
})
