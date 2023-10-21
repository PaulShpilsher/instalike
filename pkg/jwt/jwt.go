package jwt

import (
	"crypto/rsa"
	"log"
	"os"
	"time"

	"github.com/PaulShpilsher/instalike/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

type JwtSettings struct {
	Ttl        time.Duration
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewJwtSettings(config *config.JwtConfig) JwtSettings {
	return JwtSettings{
		Ttl:        time.Duration(config.TokenExpirationMinutes) * time.Minute,
		PrivateKey: getPrivateKey(config.PrivateKeyFile),
		PublicKey:  getPublicKey(config.PublicKeyFile),
	}
}

func getPrivateKey(filename string) *rsa.PrivateKey {

	key, err := jwt.ParseRSAPrivateKeyFromPEM(readFile(filename))
	if err != nil {
		log.Panicf("parsing private key failed. err: %v", err)
	}

	return key
}

func getPublicKey(filename string) *rsa.PublicKey {
	key, err := jwt.ParseRSAPublicKeyFromPEM(readFile(filename))
	if err != nil {
		log.Panicf("parsing public key failed. err: %v", err)
	}

	return key
}

func readFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Panicf("read %s files failed. err: %w", filename, err)
	}
	return data
}
