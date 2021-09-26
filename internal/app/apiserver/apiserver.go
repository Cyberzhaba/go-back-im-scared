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
	s.router.GET("/ping", s.Ping())
	s.router.GET("/user", s.GetUserByID())
	s.router.POST("/user/create", s.CreateUser())
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
