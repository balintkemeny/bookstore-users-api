package users

import (
	"github.com/balintkemeny/bookstore-users-api/domain/users"
	"github.com/balintkemeny/bookstore-users-api/services"
	errors2 "github.com/balintkemeny/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User

	/*
	bytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		//TODO: Handle the error
		return
	}

	if err = json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		//TODO: Handle the unmarshal error
		return
	}
	 */

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors2.NewBadRequestError("invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors2.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)

		return
	}

	user, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}
