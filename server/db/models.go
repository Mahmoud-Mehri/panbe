package db

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	ID uint `gorm:"primarykey"`
	PrivateKey string
	PublicKey string
	UserID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	ID uint `gorm:"primarykey"`
	Username string
	Password string
	wallets []Wallet 
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}