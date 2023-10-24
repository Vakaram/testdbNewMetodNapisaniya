package apiserver

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"testdbNewMetods/store"
	"time"
)

type Config struct {
	Env      string `yaml:"env"    env-default:"local"` // для сервера
	Port     string `yaml:"port"   env-default:":8080"`
	Server   `yaml:"server"`
	LogLevel string `yaml:"log-level"    env-default:"debug"` // для сервера
	Servis   servis
	Store    *store.Config
	//Store    *store.Config
}

type servis interface {
	SayHello() string
	TestApi() string
	CreateEmployee() string
}

type Server struct {
	Address     string        `yaml:"address"      env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout"      env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
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
	//дефолт для конфига к бд // типо передали строку с подключением
	fmt.Printf("ДО Вижу что нам вернет функция newConfig в файле Config %s \n", cfg.Store)
	cfg.Store = store.NewConfig()
	fmt.Printf("ПОСЛЕ Вижу что нам вернет функция newConfig в файле Config %s \n", cfg.Store)

	//cfg.LogLevel = "debug"
	return &cfg
}
