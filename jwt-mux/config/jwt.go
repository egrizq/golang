package config

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte(os.Getenv("JWT"))

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}
