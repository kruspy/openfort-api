package rsa_util

import (
	"crypto/rand"
	"crypto/rsa"
	"openfort-api/cmd/openfort-api/logger"
)

// GenerateKeyPair generates a new key pair.
func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		logger.Error(err)
	}
	return privKey, &privKey.PublicKey
}
