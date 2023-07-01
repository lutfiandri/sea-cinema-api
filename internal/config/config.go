package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
