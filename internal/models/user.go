package models

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" binding:"required" db:"username"`
	Email    string `json:"email" binding:"required" db:"email"`
	Password string `json:"password" binding:"required" db:"password"`
	IsAdmin  bool   `json:"is_admin" db:"is_admin"`
}
