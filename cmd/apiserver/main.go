package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Cyberzhaba/go-back-im-scared/internal/app/apiserver"
	_ "github.com/Cyberzhaba/go-back-im-scared/nftdev"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

// My code is comments :)
func main() {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalf("%v", err)
	}
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatalf("%v", err)
	}
	// nftdev.Dev()
}
