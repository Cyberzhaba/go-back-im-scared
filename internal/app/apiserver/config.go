package apiserver

import "github.com/Cyberzhaba/go-back-im-scared/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// New Config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":7777",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
