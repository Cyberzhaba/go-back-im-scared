package store

import (
	"gorm.io/gorm"
)

type UserBid struct {
	gorm.Model
	MaxValue  int    `json:"maxvaule"`
	Timestamp string `json:"timestamp"`
	TokenID   string `json:"token_id"`
	Contract  string `json:"contract"`
}

type CheckUserBid struct {
	MaxValue  int    `json:"maxvaule"`
	Timestamp string `json:"timestamp"`
	TokenID   string `json:"token_id"`
	Contract  string `json:"contract"`
}

type UserHistory struct {
	gorm.Model
	Type string `json:"type"`
}

type User struct {
	gorm.Model
	TelegramID  string        `json:"telegram_id" gorm:"unique"`
	SeedPhrase  string        `json:"seed_phrase"`
	AddrWallet  string        `json:"addr_wallet"`
	UserBids    []UserBid     `json:"user_bids" gorm:"foreignKey:id"`
	UserHistory []UserHistory `json:"user_history" gorm:"foreignKey:id"`
}

// Struct for test
type Pong struct {
	Message string `json:"message"`
}
