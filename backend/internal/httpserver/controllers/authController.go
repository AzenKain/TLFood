package controllers

import (
	"context"
	"app_food/internal/httpserver/services"
	"app_food/pkg/dtos"
	"app_food/pkg/validator"
	"app_food/pkg/models"
	"time"
	"github.com/gofiber/fiber/v2"
)


// LoginController godoc
// @Summary User login
// @Description Login a user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param dto body dtos.LoginDto true "Login data"
// @Success 200 {object} models.AuthResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /auth/login [post]
func LoginController(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    dto := &dtos.LoginDto{}

    if err := validator.ValidateDto(c, dto); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
	var res *models.AuthResponse
    res, err := services.LoginService(ctx, dto) 
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.Status(fiber.StatusOK).JSON(res)
}



// SignUpController godoc
// @Summary User registration
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param dto body dtos.SignUpDto true "Sign Up data"
// @Success 201 {object} models.AuthResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /auth/signup [post]
func SignUpController(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    dto := &dtos.SignUpDto{}

    if err := validator.ValidateDto(c, dto); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
	var res *models.AuthResponse
    res, err := services.SignUpService(ctx, dto) 
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.Status(fiber.StatusCreated).JSON(res)
}


// RefreshController godoc
// @Summary Refresh JWT token
// @Description Refresh access token using refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 201 {object} models.AuthResponse "Access token refreshed successfully"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /auth/refresh [post]
func RefreshController(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var res *models.AuthResponse
	res, err := services.RefreshService(ctx, c) 
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}


