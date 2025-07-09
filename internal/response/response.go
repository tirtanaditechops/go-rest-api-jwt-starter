package response

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(c *fiber.Ctx, status int, success bool, message string, data interface{}) error {
	return c.Status(status).JSON(ApiResponse{
		Code:    status,
		Success: success,
		Message: message,
		Data:    data,
	})
}

func Success(c *fiber.Ctx, message string, data interface{}) error {
	return JSON(c, fiber.StatusOK, true, message, data)
}

func Error(c *fiber.Ctx, status int, message string) error {
	log.Printf("API Error [%d]: %s", status, message)
	return JSON(c, status, false, message, nil)
}
