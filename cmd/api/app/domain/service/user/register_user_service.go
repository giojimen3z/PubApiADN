package user

import (
	"fmt"
	"net/http"

	"github.com/PubApiADN/cmd/api/app/domain/model"
	"github.com/PubApiADN/cmd/api/app/domain/port"
	"github.com/PubApiADN/pkg/apierrors"
	"github.com/PubApiADN/pkg/logger"
)

const (
	errorIDExist          = "User id:%v already exists"
	logErrorInvalidUserID = "User id:%v already exists  [Class: RegisterUserService][Method:RegisterUser]"
)

type RegisterUserService interface {
	// RegisterUser Send to repository the user
	RegisterUser(user model.User) apierrors.ApiError
}

type RegisterUser struct {
	UserRepository port.UserRepository
}

func (registerUser *RegisterUser) RegisterUser(user model.User) apierrors.ApiError {

	errorRepository := registerUser.UserRepository.Save(user)

	if errorRepository != nil {
		logger.Error(fmt.Sprintf(logErrorInvalidUserID, user.UserID), errorRepository)
		err := apierrors.NewApiError(fmt.Sprintf(errorIDExist, user.UserID), http.StatusText(http.StatusConflict), http.StatusConflict, nil)
		return err
	}

	return nil

}
