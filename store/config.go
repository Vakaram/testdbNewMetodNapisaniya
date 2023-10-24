package store

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	DatabaseURL string `yaml:"database_url"` // для сервера

}

func NewConfig() *Config {
	configPath := "internal/config/local.yaml" // os.Getenv("CONFIG_PATH")
	// check if file exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	// проверяю что удалось считать строку из файла
	fmt.Println(cfg.DatabaseURL)
	return &cfg

}
