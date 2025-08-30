package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL    string
	AppID      string
	AppSecret  string
	UserID     string
	UserSecret string
	Port       string
}

var Cfg *Config

func Init() {
	_ = godotenv.Load()
	Cfg = &Config{
		BaseURL:    os.Getenv("BASE_URL"),
		AppID:      os.Getenv("APP_ID"),
		AppSecret:  os.Getenv("APP_SECRET"),
		UserID:     os.Getenv("USER_ID"),
		UserSecret: os.Getenv("USER_SECRET"),
		Port:       os.Getenv("PORT"),
	}
	if Cfg.BaseURL == "" || Cfg.AppID == "" || Cfg.AppSecret == "" {
		log.Fatal("missing required envs")
	}
}
