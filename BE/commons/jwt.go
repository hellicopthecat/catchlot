package commons

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	Email    string `json:"email"`
	Social   string `json:"social"`
	Nickname string `json:"nickname"`

	jwt.RegisteredClaims
}

func GenerateACCESSJWT(email string, social string, nickname string) (string, error) {
	signKey := os.Getenv("SECRET_ACCESS_JWT_KEY")

	claims := JwtClaims{
		Email:    email,
		Social:   social,
		Nickname: nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "catch-lot-access",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	at, err := token.SignedString([]byte(signKey))
	return at, err
}

func GenerateREFRESHJWT() (string, error) {
	signKey := os.Getenv("SECRET_REFRESH_JWT_KEY")

	claims := jwt.RegisteredClaims{
		Issuer:    "catch-lot-refresh",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Subject:   "catch-lot-refresh",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rt, err := token.SignedString([]byte(signKey))
	return rt, err
}
