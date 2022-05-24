package builder

import "github.com/PubApiADN/cmd/api/app/domain/model"

type UserDataBuilder struct {
	userID   int64
	name     string
	email    string
	password string
}

func NewUserDataBuilder() *UserDataBuilder {
	return &UserDataBuilder{
		userID:   1,
		name:     "Gio",
		email:    "Gio@test.com",
		password: "asddasd",
	}
}

func (builder *UserDataBuilder) Build() model.User {
	return model.User{
		UserID:   builder.userID,
		Name:     builder.name,
		Email:    builder.email,
		Password: builder.password,
	}
}
