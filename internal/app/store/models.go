package store

import (
	"gorm.io/gorm"
)

type UserBids struct {
	ID        int `json:"id" gorm:""`
	Timestamp string
}

type UserHistory struct {
	gorm.Model
	ID   int    `json:"id"`
	Type string `json:"type"`
}

// Struct for test
type Pong struct {
	Message string `json:"message"`
}

type User struct {
	gorm.Model
	ID          int           `json:"id"`
	TelegramID  int           `json:"telegramID"`
	SeedPhrase  string        `json:"seedPhrase"`
	AddrWallet  string        `json:"addrWallet"`
	UserBids    []UserBids    `json:"userBids" gorm:"foreignKey:ID"`
	UserHistory []UserHistory `json:"userHistory" gorm:"foreignKey:ID"`
}
