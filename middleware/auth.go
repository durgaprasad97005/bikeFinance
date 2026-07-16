package middleware

import (
	"strings"

	"github.com/durgaprasad97005/bikeFinance/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

// Auth middleware that validates access token
func NewAuthMiddleware(jwtSecret string) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Access Authorization header from request
	    authHeader := c.Get("Authorization")
	    if authHeader == "" {
	    	return utils.Error(c, fiber.StatusUnauthorized, "Missing Authorization header", "Cannot find access token")
	    }

	    // Get access token from authHeader
	    const bearerPrefix = "Bearer "
	    if !strings.HasPrefix(authHeader, bearerPrefix) {
	    	return utils.Error(c, fiber.StatusUnauthorized, "Missing Bearer keyword before token", "Invlid token schema (Must be Bearer)")
	    }

	    accessToken := strings.TrimPrefix(authHeader, bearerPrefix)
	    if accessToken == "" {
	    	return utils.Error(c, fiber.StatusUnauthorized, "Missing token in Authorization header", "Cannot find Bearer token")
	    }

	    // Parse the access token
	    token, err := jwt.Parse(accessToken, func(t *jwt.Token) (any, error) {
	    	return []byte(jwtSecret), nil
	    }, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil {
			return utils.Error(c, fiber.StatusUnauthorized, "Invalid token", err.Error())
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims == nil {
			return utils.Error(c, fiber.StatusUnauthorized, "Claims didn't exist", "Error accessing token claims")
		}

		userId, ok := claims["sub"].(string)
		if !ok || userId == "" {
			return utils.Error(c, fiber.StatusUnauthorized, "Cannot access userId from token", "Error accessing userId")
		}

		// Store the userId in request locals
		c.Locals("userId", userId)
		return c.Next()
	}
}