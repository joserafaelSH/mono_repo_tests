package models

type role string

const (
	USER  role = "USER"
	ADMIN role = "ADMIN"
)

type User struct {
	ID       string `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     role   `json:"role" gorm:"type:role" `
}
