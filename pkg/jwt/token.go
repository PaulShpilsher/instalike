package jwt

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/PaulShpilsher/instalike/pkg/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSettings JwtSettings

func Initialize(config *config.JwtConfig) {
	jwtSettings = NewJwtSettings(config)
}

func CreateJwtToken(userId int) (string, error) {
	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ID:        strconv.Itoa(userId),
		IssuedAt:  jwtTime(time.Now()),
		NotBefore: jwtTime(time.Now()),
		ExpiresAt: jwtTime(time.Now().Add(jwtSettings.Ttl)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(jwtSettings.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	log.Printf("token: %s", tokenString)
	return tokenString, nil
}

func ValidateJwtToken(token string) (jwt.MapClaims, error) {

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return jwtSettings.PublicKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}

	return claims, nil
}

func jwtTime(t time.Time) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Unix(t.Unix(), 0))
}
