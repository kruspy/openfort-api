package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"openfort-api/api"
	"openfort-api/api/models"
)

// ValidateAPIKey checks if the key passed along the request belongs to a user.
func ValidateAPIKey(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.Header.Get("X-Api-Key")

		err := db.Where("api_key = ?", key).First(&models.User{}).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				api.HandleError(c, http.StatusUnauthorized, api.ErrorResponse{Message: "Not Authorized"})
				return
			}

			api.HandleError(c, http.StatusInternalServerError, api.ErrorResponse{Message: "Error authenticating API key"})
		}

		return
	}
}
