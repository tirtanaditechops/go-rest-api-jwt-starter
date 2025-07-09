package middleware

import (
	"strings"

	"loketnadi-be-go/internal/service"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if !strings.HasPrefix(token, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "format token salah"})
		}
		token = strings.TrimPrefix(token, "Bearer ")

		claims, err := service.ParseToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
		}

		// Simpan data ke context
		c.Locals("id", claims["id"])
		c.Locals("name", claims["name"])
		c.Locals("email", claims["email"])
		c.Locals("kocab", claims["kocab"])
		c.Locals("role_id", claims["role_id"])
		c.Locals("is_active", claims["is_active"])

		return c.Next()
	}
}
