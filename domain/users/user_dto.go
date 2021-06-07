package users

import (
	errors2 "github.com/balintkemeny/bookstore-users-api/utils/errors"
	"strings"
)

type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (u *User) Validate() *errors2.RestErr {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))

	if u.Email == "" {
		return errors2.NewBadRequestError("invalid email address")
	}

	return nil
}
