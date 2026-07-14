package handlers

import (
	"context"
	"time"

	"github.com/durgaprasad97005/bikeFinance/dto/request"
	"github.com/durgaprasad97005/bikeFinance/services"
	"github.com/durgaprasad97005/bikeFinance/utils"
	"github.com/gofiber/fiber/v3"
)

// struct for AuthHandler
type AuthHandler struct {
	authService *services.AuthService
}

// Constructor like func to initialize AuthHandler
func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

// Sign up controller
func (h *AuthHandler) Register(c fiber.Ctx) error {
	// Parse the request body
	var req request.RegisterUser
	if err := c.Bind().Body(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Unable to parse request body", err.Error())
	}

	// Create context for calling service
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	res, err := h.authService.Register(ctx, req) // res is response.User type
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "User creation failed with an error", err.Error())
	}

	// Success response
	return utils.Success(c, fiber.StatusCreated, "User registration successful", res)
}

// Login controller
func (h *AuthHandler) Login(c fiber.Ctx) error {
	// Parse the request body
	var req request.LoginUser
	if err := c.Bind().Body(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Unable to parse request body", err.Error())
	}

	// Create context for calling service
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	accessToken, err := h.authService.Login(ctx, req)
	if err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, "Unauthorized", err.Error())
	}

	// Success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true, 
		"message": "Login successful", 
		"accessToken": accessToken, 
	})
}

// Logout controller
func (h *AuthHandler) Logout(c fiber.Ctx) error {
	// Pending - Need to implement this function
	return nil
}

// Get profile
func (h *AuthHandler) Profile(c fiber.Ctx) error {
	// Pending - Need to implement this function
	return nil
}