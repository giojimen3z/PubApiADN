package port

import "github.com/PubApiADN/cmd/api/app/domain/model"

// UserRepository  use for all transactions about user
type UserRepository interface {
	//Save persist the user data
	Save(user model.User) (err error)
}
