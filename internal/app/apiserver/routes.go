package apiserver

import (
	"net/http"

	"github.com/Cyberzhaba/go-back-im-scared/internal/app/store"
	nft "github.com/Cyberzhaba/go-back-im-scared/nft"
	"github.com/gin-gonic/gin"
)

// Test func, return 200 and {"message" : "pong"}
func (s *APIserver) Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		pong := store.Pong{Message: "pong"}
		// data, err := json.Marshal(pong)
		// if err != nil {
		// 	return
		// }
		c.JSON(http.StatusOK, pong)
	}
}

// Get user by unique telegram-id from bot
func (s *APIserver) GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		tgid := c.Param("telegram_id")
		var user store.User
		s.store.Database.First(&user, "telegram_id = ?", tgid)
		c.JSON(http.StatusOK, user)
	}
}

// Create new user if not exists with new wallet
func (s *APIserver) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		type checkUser struct {
			TelegramID int `json:"telegram_id"`
		}

		var u checkUser
		if err := c.BindJSON(&u); err != nil {
			s.logger.Error(err)
		}
		// Generate "new" account
		account, _ := nft.GenerateWallet()
		// Create new user
		newUser := store.User{
			TelegramID: u.TelegramID,
			// TODO REPLACE TO AUTOGENERATION
			SeedPhrase: "tag volcano eight thank tide danger coast health above argue embrace heavy",
			AddrWallet: account.Address.Hex(),
		}

		// s.store.Database.Clauses(clause.OnConflict{
		// 	Columns:   []clause.Column{{Name: "telegram_id"}},

		// }).Create(&newUser)
		if err := s.store.Database.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusConflict, store.User{})
			s.logger.Error("%v", err)
			return
		}
		// DEV TEST
		// data, err := os.ReadFile("testdata.txt")
		// if err != nil {
		// 	log.Printf("%v", err)
		// }
		// nft.SignTransaction(*wallet, account, data)

		c.JSON(http.StatusCreated, newUser)
	}
}
