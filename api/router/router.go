package router

import (
	"github.com/gin-gonic/gin"
	"openfort-api/api/handlers"
	"openfort-api/api/models"
	"openfort-api/api/router/middleware"
)

// InitRouter initializes routing information and starts the database connection.
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	models.ConnectStore()

	r.POST("/encrypt", middleware.ValidateAPIKey(models.GetDatabase()), handlers.HandleEncrypt)
	r.POST("/decrypt", middleware.ValidateAPIKey(models.GetDatabase()), handlers.HandleDecrypt)

	return r
}
