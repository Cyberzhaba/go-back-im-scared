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

func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: gin.Default(),
	}
}

func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting API server...")

	return http.ListenAndServe(
		s.config.BindAddr, s.router)
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIserver) configureRouter() {
	s.router.GET("/ping", s.Ping())
	s.router.POST("/createuser", s.CreateUser())
}

func (s *APIserver) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

// func (s *APIserver) Ping() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.JSON(http.StatusOK, "OK")
// 	}
// }
