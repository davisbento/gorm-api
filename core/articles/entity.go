package articles

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model

	Id          uuid.UUID `gorm:"primaryKey;type:uuid;"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
}
