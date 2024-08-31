package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string `json:"database_url"`
	Port        string `json:"port"`
	JWTSecret   string `json:"jwt_secret"`
}

func LoadConfig(file string) (*Config, error) {
	// .envファイルから環境変数を読み込む
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var config Config
	if err := json.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}

	// 環境変数からJWT_SECRETを取得し、Configに設定
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}
	config.JWTSecret = secret

	return &config, nil
}
