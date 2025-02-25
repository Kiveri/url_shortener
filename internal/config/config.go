package config

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseType string
	BaseUrl      string
	TgBotToken   string
}

func NewConfig() *Config {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if err = godotenv.Load(filepath.Join(pwd, "../.env")); err != nil {
		log.Println("failed to load .env file")
	}
	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://localhost:8080/"
	}

	tgBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if tgBotToken == "" {
		panic("TELEGRAM_BOT_TOKEN environment variable not set")
	}

	storageType := flag.String("storage", "postgresql", "Тип хранилища (in-memory или postgresql)")
	flag.Parse()

	if *storageType != "postgresql" && *storageType != "in-memory" {
		panic("only postgresql and in-memory storages are supported")
	}
	return &Config{
		DatabaseType: *storageType,
		BaseUrl:      baseUrl,
		TgBotToken:   tgBotToken,
	}

}
