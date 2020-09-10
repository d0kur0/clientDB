package dbWrapper

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID             int    `json:"id"`
	Name           string `json:"name" validate:"required,min=3,max=36"`
	Email          string `json:"email" validate:"required,email"`
	EmailConfirmed bool   `json:"emailConfirmed" validate:"eqcsfield=Email"`
	Password       string `json:"password" validate:"min=3"`
	Avatar         string `json:"avatar"`
}
