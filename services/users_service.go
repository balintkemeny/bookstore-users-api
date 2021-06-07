package services

import (
	"github.com/balintkemeny/bookstore-users-api/domain/users"
	errors2 "github.com/balintkemeny/bookstore-users-api/utils/errors"
)

func GetUser(userID int64) (*users.User, *errors2.RestErr) {
	res := &users.User{Id: userID}

	if err := res.Get(); err != nil {
		return nil, err
	}

	return res, nil
}

func CreateUser(user users.User) (*users.User, *errors2.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
