package model

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"-"`
	Kocab        string `json:"kocab"`
	RoleID       int    `json:"role_id"`
	IsActive     bool   `json:"is_active"`
	RefreshToken string `json:"refresh_token"`
}
