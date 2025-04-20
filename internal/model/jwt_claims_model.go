package model

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
