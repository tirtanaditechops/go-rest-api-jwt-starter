package app

import (
	"loketnadi-be-go/internal/handler"
	"loketnadi-be-go/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup() *fiber.App {
	app := fiber.New()

	// Group prefix /api/auth
	auth := app.Group("/api/auth")

	// Health check
	auth.Get("/ping", handler.Ping)

	// auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
	auth.Post("/refresh", handler.RefreshToken)
	auth.Get("/me", middleware.JWTProtected(), handler.Me)
	auth.Post("/logout", middleware.JWTProtected(), handler.Logout)
	auth.Post("/register", handler.Register)

	return app
}
