package app

import (
	"github.com/balintkemeny/bookstore-users-api/controllers/health"
	"github.com/balintkemeny/bookstore-users-api/controllers/users"
)

func mapURLs() {
	router.GET("/health", health.Ping)

	router.GET("users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}