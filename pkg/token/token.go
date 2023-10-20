package token

import (
	"crypto/rsa"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	tokenExpireDuration time.Duration = time.Hour * time.Duration(24) // default is 1 hour
	privateKey          *rsa.PrivateKey
	publicKey           *rsa.PublicKey
)

func init() {
	if value, ok := os.LookupEnv("JWT_EXPIRATION_HOURS"); ok {
		if hours, err := strconv.Atoi(value); err == nil {
			tokenExpireDuration = time.Duration(hours) * time.Hour
		}
	}

	rsa, err := os.ReadFile("keys/rsa")
	if err != nil {
		log.Fatalln(err)
	}

	prvKey, err := jwt.ParseRSAPrivateKeyFromPEM(rsa)
	if err != nil {
		log.Fatalln(err)
	}

	rsaPub, err := os.ReadFile("keys/rsa.pub")
	if err != nil {
		log.Fatalln(err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(rsaPub)
	if err != nil {
		log.Fatalln(err)
	}

	privateKey = prvKey
	publicKey = pubKey
}

func CreateJwtToken(userId int) (string, error) {
	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ID:        strconv.Itoa(userId),
		IssuedAt:  jwtTime(time.Now()),
		NotBefore: jwtTime(time.Now()),
		ExpiresAt: jwtTime(time.Now().Add(tokenExpireDuration)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	log.Printf("token: %s", tokenString)
	return tokenString, nil
}

func jwtTime(t time.Time) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Unix(t.Unix(), 0))
}
