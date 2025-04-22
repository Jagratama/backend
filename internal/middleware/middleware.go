package middleware

import (
	"errors"
	"jagratama-backend/internal/helpers"
	"jagratama-backend/internal/model"
	"net/http"
	"strings"

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
		c.Set("userRole", claims.Role)

		// Continue to the next handler
		return next(c)
	}
}

func RoleCheck(roles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract the user from the context (Assuming user is set in context after JWT verification)
			userRole, ok := c.Get("userRole").(string)
			if !ok {
				return helpers.SendResponseHTTP(c, http.StatusInternalServerError, "User not valid", nil)
			}

			// Check if the user has the required role
			for _, role := range roles {
				if strings.EqualFold(userRole, role) {
					// User has the required role, continue to the next handler
					return next(c)
				}
			}

			// If the user does not have the required role, return an error
			return helpers.SendResponseHTTP(c, http.StatusForbidden, "You don't have permission to access this resource", nil)
		}
	}
}
