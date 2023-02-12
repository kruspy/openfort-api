package api

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"error"`
}

func HandleError(c *gin.Context, status int, error ErrorResponse) {
	c.JSON(status, error)
	c.Abort()
}
