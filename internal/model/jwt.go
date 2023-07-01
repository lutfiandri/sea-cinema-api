package model

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTUserClaimsData struct {
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
}

type JWTClaims struct {
	jwt.StandardClaims
	JWTUserClaimsData
}
