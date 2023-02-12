package handlers_test

import (
	"crypto/rsa"
	"openfort-api/api/handlers"
	rsa_util "openfort-api/pkg/rsa-util"
	"testing"
)

const (
	EncodedMessageLenght int = 684
)

func TestHandlerEncrypt_Encrypt(t *testing.T) {
	var publicKey *rsa.PublicKey

	// Generate the key pair for the test.
	_, publicKey = rsa_util.GenerateKeyPair(4096)

	// Convert the public key to bytes.
	pubBytes, _ := rsa_util.PublicKeyToBytes(publicKey)

	tests := []struct {
		name    string
		key     []byte
		message string
		expLen  int
	}{
		{
			name:    "Encrypted message length",
			key:     pubBytes,
			message: "Secret message",
			expLen:  EncodedMessageLenght,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc, _ := handlers.Encrypt(tt.key, tt.message)
			if len(enc) != tt.expLen {
				t.Errorf("wanted %d, got %d", tt.expLen, len(enc))
			}
		})
	}
}
