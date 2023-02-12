package models

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"openfort-api/cmd/openfort-api/config"
	"openfort-api/cmd/openfort-api/logger"
	aes_util "openfort-api/pkg/aes-util"
	rsa_util "openfort-api/pkg/rsa-util"
	"os"
)

var db *gorm.DB

// ConnectStore initializes the database instance by starting a connection to the database.
func ConnectStore() {
	cfg := config.GetConfig()

	var err error
	db, err = gorm.Open(postgres.Open(cfg.DbUrl), &gorm.Config{})

	if err != nil {
		logger.Fatal(fmt.Sprintf("[DB] Error openning connection to database: %s", err.Error()))
	}

	err = db.AutoMigrate(&User{}, &Keys{})
	if err != nil {
		logger.Fatal(fmt.Sprintf("[DB] Error migrating database: %s", err.Error()))
	}

	createMockKeys() // This is temporary. Will be removed once the API has endpoints to create users and keys.
}

// GetDatabase returns the current database instance.
func GetDatabase() *gorm.DB {
	return db
}

type MockKey struct {
	Key string `json:"key"`
	Pwd string `json:"pwd"`
}
type MockKeys struct {
	Keys []MockKey `json:"keys"`
}

// createMockKeys creates sample data to interact with the API.
func createMockKeys() {
	data, err := os.ReadFile("./conf/mock.json")
	if err != nil {
		logger.Error(err)
	}

	mockKeys := &MockKeys{}
	err = json.Unmarshal(data, mockKeys)
	if err != nil {
		logger.Error(err)
	}

	for _, mockKey := range mockKeys.Keys {
		createSingleMockKey(mockKey.Key, mockKey.Pwd)
	}
}

// createSingleMockKey creates a user and keys mock and saves it to the database.
func createSingleMockKey(key, pwd string) {
	if db.FirstOrCreate(&User{ApiKey: key}).RowsAffected == 0 {
		return
	}

	priv, pub := rsa_util.GenerateKeyPair(4096)

	pubBytes, _ := rsa_util.PublicKeyToBytes(pub)
	privBytes := rsa_util.PrivateKeyToBytes(priv)

	encPrivBytes, _ := aes_util.Encrypt(privBytes, pwd)

	db.FirstOrCreate(&Keys{
		ApiKey:     key,
		PublicKey:  pubBytes,
		PrivateKey: encPrivBytes,
	})
}
