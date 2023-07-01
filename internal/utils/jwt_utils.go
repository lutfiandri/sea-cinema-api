package utils

import (
	"time"

	"sea-cinema-api/internal/config"
	"sea-cinema-api/internal/model"

	"github.com/golang-jwt/jwt"
)

func GenerateJwt(userClaimsData model.JWTUserClaimsData) (string, error) {
	claims := model.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
		JWTUserClaimsData: userClaimsData,
	}

	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokens.SignedString([]byte(config.JWTSecretKey))
}

func ParseJwt(tokenString string) (model.JWTClaims, error) {
	var claims model.JWTClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecretKey), nil
	})
	if err != nil || !token.Valid {
		return claims, err
	}

	return claims, err
}
