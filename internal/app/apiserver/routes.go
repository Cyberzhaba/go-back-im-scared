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
		tgid := c.Query("telegram_id")
		s.logger.Debug(tgid)
		var user store.User
		s.store.Database.First(&user, "telegram_id = ?", tgid)
		c.JSON(http.StatusOK, user)
	}
}

// Create new user if not exists with new wallet
func (s *APIserver) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// VALIDATE ON FRONTEND

		type checkUser struct {
			TelegramID string `json:"telegram_id"`
		}

		var u checkUser
		if err := c.BindJSON(&u); err != nil {
			s.logger.Error(err)
			c.JSON(http.StatusConflict, store.User{})
			return
		}

		tgid := u.TelegramID
		if tgid == "" {
			s.logger.Error("telegram-id is empty")
			c.JSON(http.StatusConflict, store.User{})
			return
		}
		// Generate "new" account
		account, _ := nft.GenerateWallet()
		// Create new user
		newUser := store.User{
			TelegramID: tgid,
			// TODO REPLACE TO AUTOGENERATION
			SeedPhrase: "tag volcano eight thank tide danger coast health above argue embrace heavy",
			AddrWallet: account.Address.Hex(),
		}

		if err := s.store.Database.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusConflict, store.User{})
			s.logger.Error("%v", err)
			return
		}
		// DEV TEST
		// data, err := os.ReadFile("testdata.txt")
		// if err != nil {
		// 	s.logger.Error(err)
		// }
		// s.logger.Info(data)
		// sign, err := wallet.SignDataWithPassphrase(account, "tag volcano eight thank tide danger coast health above argue embrace heavy", "text/plain", data)
		// fmt.Println(sign, err)
		// addrm := account.Address
		// fmt.Println(addrm, err)
		// nft.SignTransaction(*wallet, account, data)

		c.JSON(http.StatusCreated, newUser)
	}
}

//
// func (s *APIserver) GetUserBids() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var bids []store.UserBid
// 		// var user store.User
// 		s.store.Database.Model(bids)
// 		c.JSON(http.StatusOK, bids)
// 	}
// }

// Create bid
func (s *APIserver) CreateBid() gin.HandlerFunc {
	return func(c *gin.Context) {
		var checkBid store.CheckUserBid
		var bid store.UserBid
		if err := c.BindJSON(&checkBid); err != nil {
			s.logger.Error(err)
			c.JSON(http.StatusConflict, store.UserBid{})
			return
		}
		bid = store.UserBid{
			TelegramID: checkBid.TelegramID,
			MaxValue:   checkBid.MaxValue,
			Timestamp:  checkBid.Timestamp,
			TokenID:    checkBid.TokenID,
			Contract:   checkBid.Contract,
		}
		if err := s.store.Database.Create(&bid).Error; err != nil {
			c.JSON(http.StatusConflict, store.UserBid{})
			s.logger.Error("%v", err)
			return
		}
		c.JSON(http.StatusOK, bid)
	}
}
