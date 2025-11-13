package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	configPath = "config/config.yaml"
)

type Config struct {
	Server Server `yaml:"server"`
	DB     DB
}

type Server struct {
	Host         string        `yaml:"host"`
	Port         int           `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type DB struct {
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	Port     int    `env:"DB_PORT"`
	Driver   string `env:"DB_DRIVER"`
}

func InitConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	if err := cleanenv.ReadEnv(&cfg.DB); err != nil {
		return nil, err
	}
	return &cfg, nil
}
