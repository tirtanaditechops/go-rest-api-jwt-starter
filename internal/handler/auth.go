package handler

import (
	"loketnadi-be-go/internal/model"
	"loketnadi-be-go/internal/response"
	"loketnadi-be-go/internal/service"
	"loketnadi-be-go/pkg/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil || req.Email == "" || req.Password == "" {
		return response.Error(c, fiber.StatusBadRequest, "email dan password wajib diisi")
	}

	user, err := service.LoginUser(req.Email, req.Password)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, err.Error())
	}

	accessToken, err := service.GenerateJWT(user)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "gagal generate token")
	}

	refreshToken, err := service.GenerateRefreshToken(user)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "gagal generate refresh token")
	}

	return response.Success(c, "login berhasil", fiber.Map{
		"token":         accessToken,
		"refresh_token": refreshToken,
		"expires_in":    86400,
		"user": fiber.Map{
			"id":        user.ID,
			"name":      user.Name,
			"email":     user.Email,
			"kocab":     user.Kocab,
			"role_id":   user.RoleID,
			"is_active": user.IsActive,
		},
	})
}

func Me(c *fiber.Ctx) error {
	return response.Success(c, "user aktif", fiber.Map{
		"id":        c.Locals("id"),
		"name":      c.Locals("name"),
		"email":     c.Locals("email"),
		"kocab":     c.Locals("kocab"),
		"role_id":   c.Locals("role_id"),
		"is_active": c.Locals("is_active"),
	})
}

func RefreshToken(c *fiber.Ctx) error {
	type Req struct {
		RefreshToken string `json:"refresh_token"`
	}
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "format salah")
	}

	claims, err := service.ParseToken(req.RefreshToken)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "refresh token tidak valid")
	}

	userID := int(claims["id"].(float64))
	var user model.User
	err = database.DB.QueryRow(`
SELECT id, name, email, kocab, role_id, is_active, refresh_token
FROM users WHERE id = @p1
`, userID).Scan(&user.ID, &user.Name, &user.Email, &user.Kocab, &user.RoleID, &user.IsActive, &user.RefreshToken)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "pengguna tidak ditemukan")
	}

	if user.RefreshToken != req.RefreshToken {
		return response.Error(c, fiber.StatusUnauthorized, "refresh token tidak cocok")
	}

	newAccessToken, err := service.GenerateJWT(user)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "gagal buat token baru")
	}
	newRefreshToken, err := service.GenerateRefreshToken(user)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "gagal buat refresh token baru")
	}

	return response.Success(c, "refresh berhasil", fiber.Map{
		"token":         newAccessToken,
		"refresh_token": newRefreshToken,
	})
}

func Logout(c *fiber.Ctx) error {
	userID := int(c.Locals("id").(float64))
	_, err := database.DB.Exec("UPDATE users SET refresh_token = NULL WHERE id = @p1", userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "gagal logout")
	}
	return response.Success(c, "logout berhasil", nil)
}

func Register(c *fiber.Ctx) error {
	var req struct {
		NIPP     string `json:"nipp"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Kocab    string `json:"kocab"`
		RoleID   int    `json:"role_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "format input salah")
	}

	var exists int
	err := database.DB.QueryRow("SELECT COUNT(1) FROM users WHERE email = @p1", req.Email).Scan(&exists)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "gagal cek email")
	}
	if exists > 0 {
		return response.Error(c, fiber.StatusBadRequest, "email / username sudah terdaftar")
	}

	hashedPassword, err := service.HashPassword(req.Password)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "gagal hash password")
	}

	now := time.Now()
	_, err = database.DB.Exec(`
INSERT INTO users (nipp, name, email, password, kocab, role_id, is_active, created_at, updated_at)
VALUES (@p1, @p2, @p3, @p4, @p5, @p6, 1, @p7, @p8)
`, req.NIPP, req.Name, req.Email, hashedPassword, req.Kocab, req.RoleID, now, now)

	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "gagal simpan user")
	}

	return response.Success(c, "registrasi berhasil", nil)
}
