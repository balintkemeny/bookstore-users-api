package users

import (
	"fmt"
	errors2 "github.com/balintkemeny/bookstore-users-api/utils/errors"
	"github.com/balintkemeny/bookstore-users-api/utils/time_utils"
)

var usersDB = make(map[int64]*User)

func (u *User) Get() *errors2.RestErr {
	res, ok := usersDB[u.Id]
	if !ok {
		return errors2.NewNotFoundError("user not found with the provided id")
	}

	u.Id = res.Id
	u.FirstName = res.FirstName
	u.LastName = res.LastName
	u.Email = res.Email
	u.DateCreated = res.DateCreated

	return nil
}

func (u *User) Save() *errors2.RestErr {
	_, ok := usersDB[u.Id]
	if ok {
		return errors2.NewBadRequestError(fmt.Sprintf("user %d already exists", u.Id))
	}

	u.DateCreated = time_utils.GetNowString()
	usersDB[u.Id] = u

	return nil
}