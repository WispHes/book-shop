package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database   Database
	ServerPort string `envconfig:"SERVER_PORT" default:"8080"`
}

type Database struct {
	Host     string `envconfig:"POSTGRES_HOST" required:"true"`
	Port     string `envconfig:"POSTGRES_PORT" required:"true"`
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	DBName   string `envconfig:"POSTGRES_DB" required:"true"`
	SSLMode  string `envconfig:"SSL_MODE" required:"true"`
}

func NewParsedConfig() (*Config, error) {
	_ = godotenv.Load(".env")
	cnf := Config{}
	err := envconfig.Process("", &cnf)
	return &cnf, err
}
