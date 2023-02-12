package handlers

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"openfort-api/api"
	"openfort-api/api/models"
	aes_util "openfort-api/pkg/aes-util"
	rsa_util "openfort-api/pkg/rsa-util"
)

type decryptRequestsBody struct {
	Message  string `json:"message"  binding:"required"`
	Password string `json:"password" binding:"required"`
}

type decryptResponseBody struct {
	Message string `json:"message"`
}

// HandleDecrypt deals with the POST /decrypt request.
func HandleDecrypt(c *gin.Context) {
	body := &decryptRequestsBody{}

	if err := c.BindJSON(body); err != nil {
		api.HandleError(c, http.StatusBadRequest, api.ErrorResponse{Message: "Invalid request JSON format"})
		return
	}

	key, err := models.GetPrivateKey(c.Request.Header.Get("X-Api-Key"))
	if err != nil {
		api.HandleError(c, http.StatusInternalServerError, api.ErrorResponse{Message: "Error decrypting message"})
		return
	}

	msg, err := Decrypt(key, body.Message, body.Password)
	if err != nil {
		api.HandleError(c, http.StatusInternalServerError, api.ErrorResponse{Message: "Error decrypting message"})
		return
	}

	c.JSON(http.StatusOK, &decryptResponseBody{Message: msg})
}

// Decrypt a message using the caller private key.
func Decrypt(cypher []byte, encMsg, pwd string) (string, error) {
	msg, err := base64.StdEncoding.DecodeString(encMsg)
	if err != nil {
		return "", err
	}

	bytes, err := aes_util.Decrypt(cypher, pwd)
	if err != nil {
		return "", err
	}

	priv, err := rsa_util.BytesToPrivateKey(bytes)
	if err != nil {
		return "", err
	}

	dec, err := rsa_util.DecryptWithPrivateKey(msg, priv)
	if err != nil {
		return "", err
	}

	return string(dec), nil
}
