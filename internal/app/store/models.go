package store

import (
	"gorm.io/gorm"
)

type UserBids struct {
	gorm.Model
	Timestamp string
}

type UserHistory struct {
	gorm.Model
	Type string `json:"type"`
}

type User struct {
	gorm.Model
	TelegramID  int           `json:"telegram_id" gorm:"unique"`
	SeedPhrase  string        `json:"seed_phrase"`
	AddrWallet  string        `json:"addr_wallet"`
	UserBids    []UserBids    `json:"user_bids" gorm:"foreignKey:id"`
	UserHistory []UserHistory `json:"user_history" gorm:"foreignKey:id"`
}

// Struct for test
type Pong struct {
	Message string `json:"message"`
}
