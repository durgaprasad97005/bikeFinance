package handlers

import (
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

	res, err := h.authService.Register(req) // res is response.User type
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

	accessToken, err := h.authService.Login(req)
	if err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, "Unauthorized", err.Error())
	}

	// Success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":     true,
		"message":     "Login successful",
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
	// Get the id from locals
	id := c.Locals("userId").(string)

	// Get user
	res, err := h.authService.Profile(id)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Error finding user", err.Error())
	}
	if res == nil {
		return utils.Error(c, fiber.StatusNotFound, "User not found", "Invalid Id")
	}

	// Success response
	return utils.Success(c, fiber.StatusOK, "Successfully retrieved user profile", res)
}
