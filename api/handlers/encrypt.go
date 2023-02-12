package handlers

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"openfort-api/api"
	"openfort-api/api/models"
	rsa_util "openfort-api/pkg/rsa-util"
)

type encryptRequestBody struct {
	Message string `json:"message" binding:"required"`
}

type encryptResponseBody struct {
	Message string `json:"message"`
}

// HandleEncrypt deals with the POST /encrypt request.
func HandleEncrypt(c *gin.Context) {
	body := &encryptRequestBody{}

	if err := c.BindJSON(body); err != nil {
		api.HandleError(c, http.StatusBadRequest, api.ErrorResponse{Message: "Invalid request JSON format"})
		return
	}

	key, err := models.GetPublicKey(c.Request.Header.Get("X-Api-Key"))
	if err != nil {
		api.HandleError(c, http.StatusInternalServerError, api.ErrorResponse{Message: "Error encrypting message"})
		return
	}

	msg, err := Encrypt(key, body.Message)
	if err != nil {
		api.HandleError(c, http.StatusInternalServerError, api.ErrorResponse{Message: "Error decrypting message"})
		return
	}

	c.JSON(http.StatusOK, &encryptResponseBody{Message: msg})
}

// Encrypt a message using the caller public key.
func Encrypt(pubKey []byte, msg string) (string, error) {
	pub, err := rsa_util.BytesToPublicKey(pubKey)
	if err != nil {
		return "", err
	}

	enc, err := rsa_util.EncryptWithPublicKey([]byte(msg), pub)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(enc), nil
}
