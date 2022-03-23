package users

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Id           uuid.UUID `gorm:"primaryKey;type:uuid;"`
	Name         string    `gorm:"size:100"`
	Email        string    `gorm:"size:100"`
	Password     string    `gorm:"size:100"`
	Activate     bool      `gorm:"default:true"`
	RefreshToken string    `gorm:"size:100"`
	CreatedAt    time.Time `gorm:"default:NOW()"`
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
