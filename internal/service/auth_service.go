package service

import (
	"database/sql"
	"errors"
	"os"
	"time"

	"loketnadi-be-go/internal/model"
	"loketnadi-be-go/pkg/database"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey []byte

func InitAuthService() {
	jwtKey = []byte(os.Getenv("JWT_SECRET"))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func RegisterUser(user model.User) error {
	hashed, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	_, err = database.DB.Exec("INSERT INTO users (name, email, password) VALUES (@p1, @p2, @p3)", user.Name, user.Email, hashed)
	return err
}

func LoginUser(email, password string) (model.User, error) {
	var user model.User
	err := database.DB.QueryRow(`
		SELECT id, name, email, password, kocab, role_id, is_active 
		FROM users WHERE email = @p1
	`, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Kocab, &user.RoleID, &user.IsActive)
	if err == sql.ErrNoRows {
		return user, errors.New("email tidak ditemukan")
	}
	if err != nil {
		return user, err
	}
	if !CheckPasswordHash(password, user.Password) {
		return user, errors.New("password salah")
	}
	if !user.IsActive {
		return user, errors.New("akun tidak aktif")
	}
	return user, nil
}

func GenerateJWT(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"name":      user.Name,
		"email":     user.Email,
		"kocab":     user.Kocab,
		"role_id":   user.RoleID,
		"is_active": user.IsActive,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(jwtKey)
}

func GenerateRefreshToken(user model.User) (string, error) {
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(7 * 24 * time.Hour).Unix(),
	})
	tokenString, err := refresh.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	// Simpan ke DB
	_, err = database.DB.Exec("UPDATE users SET refresh_token = @p1 WHERE id = @p2", tokenString, user.ID)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("token tidak valid")
	}
	return token.Claims.(jwt.MapClaims), nil
}
