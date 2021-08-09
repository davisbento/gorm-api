package users

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Id           uuid.UUID `gorm:"primaryKey;type:uuid;"`
	Name         string    `gorm:"size:255"`
	Email        string    `gorm:"size:255"`
	Password     string    `gorm:"size:255"`
	Activate     bool      `gorm:"default:true"`
	RefreshToken string    `gorm:"size:255"`
}

type UserCreateModel struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserList struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
