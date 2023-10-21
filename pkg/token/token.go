package token

import (
	"crypto/rsa"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/PaulShpilsher/instalike/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	Ttl        time.Duration
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewJwtService(config *config.JwtConfig) *service {
	return &service{
		Ttl:        time.Duration(config.TokenExpirationMinutes) * time.Minute,
		PrivateKey: getPrivateKey(config.PrivateKeyFile),
		PublicKey:  getPublicKey(config.PublicKeyFile),
	}
}

type JwtService interface {
	CreateToken(content string) (string, error)
	ValidateToken(token string) (string, error)
}

func (s *service) CreateToken(content string) (string, error) {

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["dat"] = content               // Our custom data.
	claims["exp"] = now.Add(s.Ttl).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()            // The time at which the token was issued.
	claims["nbf"] = now.Unix()            // The time before which the token must be disregarded.

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(s.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return token, nil
}

func (s *service) ValidateToken(token string) (string, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return s.PublicKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("token parsing failed. err: %v", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return "", fmt.Errorf("invalid token. err:  %v", err)
	}

	content := claims["dat"].(string)
	return content, nil

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
		log.Panicf("read %s files failed. err: %v", filename, err)
	}
	return data
}