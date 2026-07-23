package utils

import "github.com/gofiber/fiber/v3"

// Success response
func Success(c fiber.Ctx, status int, message string, data any) error {
	return c.Status(status).JSON(fiber.Map{
		"success": true, 
		"message": message, 
		"data": data, 
	})
}

// Error response
func Error(c fiber.Ctx, status int, message string, errs any) error {
	resultErrors := []string{}

	// Type Switch - a special switch case in Golang
	switch value := errs.(type) {
	case string:
		resultErrors = append(resultErrors, value)
	case []string:
		resultErrors = append(resultErrors, value...) // append() is a variadic function
	}

	return c.Status(status).JSON(fiber.Map{
		"success": false, 
		"message": message, 
		"errors": resultErrors, 
	})
}