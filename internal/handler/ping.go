package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"loketnadi-be-go/internal/logger"
	"loketnadi-be-go/internal/response"
)

func Ping(c *fiber.Ctx) error {
	start := time.Now()
	requestID := c.Locals("request_id")

	logger.Log.WithFields(map[string]interface{}{
		"method":     c.Method(),
		"path":       c.Path(),
		"ip":         c.IP(),
		"user_agent": c.Get("User-Agent"),
		"request_id": requestID,
	}).Info("📥 Endpoint /ping dipanggil")

	// Respon sederhana
	res := response.Success(c, "pong", nil)

	logger.Log.WithFields(map[string]interface{}{
		"request_id": requestID,
		"latency":    time.Since(start).Milliseconds(),
	}).Info("📤 Respon /ping berhasil dikirim")

	return res
}
