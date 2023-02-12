package handlers_test

import (
	"crypto/rsa"
	"openfort-api/api/handlers"
	aes_util "openfort-api/pkg/aes-util"
	rsa_util "openfort-api/pkg/rsa-util"
	"testing"
)

const (
	Message  string = "Secret message"
	Password string = "password"
)

func TestHandlerDecrypt_Decrypt(t *testing.T) {
	var privateKey *rsa.PrivateKey
	var publicKey *rsa.PublicKey

	// Generate the key pair for the test.
	privateKey, publicKey = rsa_util.GenerateKeyPair(4096)

	// Encode a message to decrypt in the test.
	pubBytes, _ := rsa_util.PublicKeyToBytes(publicKey)
	encMessage, _ := handlers.Encrypt(pubBytes, Message)

	// Encode the private key with the passphrase.
	privBytes := rsa_util.PrivateKeyToBytes(privateKey)
	encPrivKey, _ := aes_util.Encrypt(privBytes, Password)

	tests := []struct {
		name           string
		cypher         []byte
		encodedMessage string
		passphrase     string
		expMsg         string
	}{
		{
			name:           "Decrypt message",
			cypher:         encPrivKey,
			encodedMessage: encMessage,
			passphrase:     Password,
			expMsg:         Message,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dec, _ := handlers.Decrypt(tt.cypher, tt.encodedMessage, tt.passphrase)
			if dec != tt.expMsg {
				t.Errorf("wanted %s, got %s", tt.expMsg, dec)
			}
		})
	}
}
