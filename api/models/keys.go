package models

import (
	"gorm.io/gorm"
	"time"
)

type Keys struct {
	ID        uint `gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time

	ApiKey     string `json:"api_key" gorm:"primarykey"`
	PublicKey  []byte `json:"public_key"`
	PrivateKey []byte `json:"private_key"`
}

func (k *Keys) BeforeCreate(db *gorm.DB) error {
	k.CreatedAt = time.Now().Local()
	return nil
}

func (k *Keys) BeforeUpdate(db *gorm.DB) error {
	k.UpdatedAt = time.Now().Local()
	return nil
}

// GetPublicKey returns a public key given an api key.
func GetPublicKey(apiKey string) ([]byte, error) {
	keys := &Keys{}
	err := db.Select("public_key").Where(Keys{ApiKey: apiKey}).First(keys).Error
	if err != nil {
		return nil, err
	}
	return keys.PublicKey, nil
}

// GetPrivateKey returns a private key given an api key.
func GetPrivateKey(apiKey string) ([]byte, error) {
	keys := &Keys{}
	err := db.Select("private_key").Where(Keys{ApiKey: apiKey}).First(keys).Error
	if err != nil {
		return nil, err
	}
	return keys.PrivateKey, nil
}
