package token

import (
	"crypto/rsa"
	"fmt"
	"log"
	"time"

	"github.com/PaulShpilsher/instalike/pkg/config"
	"github.com/PaulShpilsher/instalike/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
)

//
// JwtService - token service implementation
//

type jwtService struct {
	Ttl        time.Duration
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewJwtService(config *config.ServerConfig) *jwtService {
	return &jwtService{
		Ttl:        time.Duration(config.TokenExpirationMinutes) * time.Minute,
		PrivateKey: getPrivateKey(config.PrivateKeyFile),
		PublicKey:  getPublicKey(config.PublicKeyFile),
	}
}

type JwtService interface {
	CreateToken(content string) (string, error)
	ValidateToken(token string) (string, error)
	TTL() time.Duration
}

func (s *jwtService) TTL() time.Duration {
	return s.Ttl
}

func (s *jwtService) CreateToken(content string) (string, error) {

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

func (s *jwtService) ValidateToken(token string) (string, error) {
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

	key, err := jwt.ParseRSAPrivateKeyFromPEM(utils.ReadFile(filename))
	if err != nil {
		log.Panicf("parsing private key failed. err: %v", err)
	}

	return key
}

func getPublicKey(filename string) *rsa.PublicKey {
	key, err := jwt.ParseRSAPublicKeyFromPEM(utils.ReadFile(filename))
	if err != nil {
		log.Panicf("parsing public key failed. err: %v", err)
	}

	return key
}
