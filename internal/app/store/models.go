package store

type UserBids struct {
	ID int `json:"id"`
}

type UserHistory struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

type Pong struct {
	Message string `json:"message"`
}

type User struct {
	TelegramID int    `json:"telegramID"`
	SeedPhrase string `json:"seedPhrase"`
	AddrWallet string `json:"addrWallet"`
}
