package middleware

import (
	"errors"
	"jagratama-backend/internal/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
		if !ok {
			return errors.New("JWT token missing or invalid")
		}
		claims, ok := token.Claims.(*model.JwtCustomClaims)
		if !ok {
			return errors.New("failed to cast claims as jwt.MapClaims")
		}

		// Set the user ID in the context
		c.Set("userID", claims.ID)

		// Continue to the next handler
		return next(c)
	}
}
