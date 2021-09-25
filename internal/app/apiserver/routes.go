package apiserver

import (
	"net/http"

	"github.com/Cyberzhaba/go-back-im-scared/internal/app/store"
	nft "github.com/Cyberzhaba/go-back-im-scared/nft"
	"github.com/gin-gonic/gin"
)

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

func (s *APIserver) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		type checkUser struct {
			TelegramID int `json:"telegramID"`
		}

		var u checkUser
		if err := c.BindJSON(&u); err != nil {
			s.logger.Error(err)
		}
		// Generate "new" account
		account := nft.GenerateWallet()
		// Create new user
		newUser := store.User{
			TelegramID: u.TelegramID,
			// TODO REPLACE TO AUTOGENERATION
			SeedPhrase: "tag volcano eight thank tide danger coast health above argue embrace heavy",
			AddrWallet: account.Address.Hex(),
		}
		s.store.Database.Create(&newUser)

		c.JSON(http.StatusCreated, newUser)
	}
}
