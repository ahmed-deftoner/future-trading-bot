package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Key struct {
	Keyid           uuid.UUID `gorm:"primary_key;type:uuid;default:gen_random_uuid()" json:"key_id"`
	ApiKey          string    `gorm:"not null;unique" json:"api_key"`
	SecretKey       string    `gorm:"not null;unique" json:"secret_key"`
	Passphrase      string    `gorm:"" json:"passphrase"`
	UserEmail       string    `json:"user_email"`
	OpenShort       int
	OpenLong        int
	TradeAmount     int
	AllowedCoins    int
	CapitalPerTrade float64
	Start           bool
}

func (u *Key) FindAllKeys(db *gorm.DB) (*[]Key, error) {
	Keys := []Key{}
	err := db.Model(&Key{}).Find(&Keys).Error
	if err != nil {
		return &[]Key{}, err
	}
	return &Keys, nil
}

func (u *Key) FindKeysByEmail(db *gorm.DB, email string) (*[]Key, error) {
	Keys := []Key{}
	err := db.Model(Key{}).Where("user_email = ?", email).Find(&Keys).Error
	if err != nil {
		return &[]Key{}, err
	}
	return &Keys, nil
}
