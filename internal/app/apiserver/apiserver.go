package apiserver

import (
	"net/http"

	"github.com/Cyberzhaba/go-back-im-scared/internal/app/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *gin.Engine
	store  *store.Store
}

// Create apiserver
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: gin.Default(),
	}
}

// Init
func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting API server on http://0.0.0.0", s.config.BindAddr)

	return http.ListenAndServe(
		s.config.BindAddr, s.router)
}

// Configure logger, set value from config file
func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

// Add routes
func (s *APIserver) configureRouter() {
	// Test func 200 {"message" : "pong"}
	s.router.GET("/ping", s.Ping())

	// Get user info
	// {"telegram_id" : string}
	s.router.GET("/get/user", s.GetUserByID())

	// Get all user's bids
	// {"telegram_id" : string}
	// s.router.GET("/get/bids", s.GetBids())

	// Create and return new user (wallet)
	s.router.POST("/create/user", s.CreateUser())

	// Create new and bid
	s.router.POST("/create/bid", s.CreateBid())

	s.router.GET("/get/bids_by_item", s.GetOrderBidsByItem())
	s.router.GET("/get/items", s.GetAllItems())
}

// Configure db, from config file
func (s *APIserver) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}
