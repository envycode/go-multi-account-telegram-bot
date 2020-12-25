package config

import (
	"github.com/subosito/gotenv"
	"log"
	"os"
)

type Config struct {
	Port int
}

func newConfig() Config {
	return Config{Port: GetInt("APP_PORT", 8080)}
}

var appConfig = Config{}

func AppConfig() Config {
	return appConfig
}

func init() {
	if err := gotenv.Load(); err != nil {
		log.Println("loading config from os environment variable")
	}
	log.SetOutput(os.Stdout)
	appConfig = newConfig()
}
