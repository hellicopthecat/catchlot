package commons

import (
	"fmt"
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

func DecodedACCESSJWT(tokenStr string) (*JwtClaims, error) {
	signKey := os.Getenv("SECRET_ACCESS_JWT_KEY")

	parseToken, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected Signing Methods")
		}
		return []byte(signKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := parseToken.Claims.(*JwtClaims)
	if !ok || !parseToken.Valid {
		return nil, fmt.Errorf("Invalid Token")
	}
	return claims, nil
}
