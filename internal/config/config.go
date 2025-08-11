package config

import (
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port   int    `env:"APP_PORT" yaml:"app_port"`
	Host   string `env:"APP_HOST" yaml:"app_host"`
	DB_Url string `env:"DB_URL" yaml:"db_url"`
}

func LoadСonfig(path string) *Config {
	config := new(Config)
	err := cleanenv.ReadConfig(path, config)
	if err != nil {
		slog.Error("Error loading config", err.Error())
	}
	err = cleanenv.ReadEnv(config)
	if err != nil {
		slog.Error("Error loading config", err.Error())
	}
	slog.Info("config loaded successfully")
	slog.Info("Config: %v", config)
	return config
}
