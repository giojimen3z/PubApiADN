package mock

import (
	"github.com/stretchr/testify/mock"

	"github.com/PubApiADN/cmd/api/app/domain/model"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (mock *UserRepositoryMock) Save(user model.User) (err error) {
	args := mock.Called(user)
	return args.Error(0)
}
