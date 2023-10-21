package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Url string
}

type JwtConfig struct {
	TokenExpirationMinutes int
	PrivateKeyFile         string
	PublicKeyFile          string
}

type DatabaseConfig struct {
	Url                   string
	MaxOpenConnections    int
	MaxIdleConnections    int
	MaxLifetimeConnctions int
}

type Config struct {
	Server   ServerConfig
	Jwt      JwtConfig
	Database DatabaseConfig
}

func LoadConfig() Config {
	loadEnv()

	return Config{

		Server: ServerConfig{
			Url: getString("SERVER_URL"),
		},

		Jwt: JwtConfig{
			TokenExpirationMinutes: getInt("TOKEN_EXPIRATION_MINUTES"),
			PrivateKeyFile:         getString("TOKEN_PRIVATE_KEY_FILE"),
			PublicKeyFile:          getString("TOKEN_PUBLIC_KEY_FILE"),
		},

		Database: DatabaseConfig{
			Url:                   getString("DB_URL"),
			MaxOpenConnections:    getInt("DB_MAX_CONNECTIONS"),
			MaxIdleConnections:    getInt("DB_MAX_IDLE_CONNECTIONS"),
			MaxLifetimeConnctions: getInt("DB_MAX_LIFETIME_CONNECTIONS"),
		},
	}
}

func loadEnv() {

	ex, err := os.Executable()
	if err != nil {
		log.Panicln(err)
	}

	envFile := filepath.Join(filepath.Dir(ex), ".env")

	if err = godotenv.Load(envFile); err != nil {
		log.Panicf("failed to load %s file. err: %v", envFile, err)
	}
}

func getString(key string) string {
	envValue, ok := os.LookupEnv(key)
	if !ok || envValue == "" {
		log.Panicf("missing environment variable %s", key)
	}
	return envValue
}

func getInt(key string) int {
	envValue, ok := os.LookupEnv(key)
	if !ok || envValue == "" {
		log.Panicf("missing environment variable %s", key)
	}

	value, err := strconv.Atoi(envValue)
	if err != nil {
		log.Panicf("unable to parse environment variable %s err: %v", key, err)
	}

	return value
}
