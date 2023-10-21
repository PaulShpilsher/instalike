package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Server struct {
	Url                string
	TokenExpireMinutes int
	PrivateKeyFile     string
	PublicKeyFile      string
}

type Database struct {
	Url                  string
	MaxOpenConnections   int
	MaxIdleConnections   int
	ConnctionMaxLifetime int
}

type Config struct {
	Server
	Database
}

func NewConfig() Config {
	return Config{}
}

func loadEnv() {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	envFile := filepath.Join(filepath.Dir(ex), ".env")

	if err = godotenv.Load(envFile); err != nil {
		panic("Error loading .env file")
	}
}
