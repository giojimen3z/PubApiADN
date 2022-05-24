package user

import (
	"github.com/PubApiADN/cmd/api/app/domain/model"
	"github.com/PubApiADN/cmd/api/app/domain/service/user"
	"github.com/PubApiADN/pkg/apierrors"
)

// RegisterUserApplication is the initial flow entry to create one user
type RegisterUserApplication interface {
	// Handler is the input for access to create one user
	Handler(user model.User) apierrors.ApiError
}

type RegisterUser struct {
	RegisterUserService user.RegisterUserService
}

func (registerUser *RegisterUser) Handler(user model.User) apierrors.ApiError {
	return registerUser.RegisterUserService.RegisterUser(user)
}
