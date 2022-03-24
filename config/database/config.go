package database

import (
	"github.com/davisbento/gorm-api/config/env"
	"github.com/davisbento/gorm-api/core/articles"
	"github.com/davisbento/gorm-api/core/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"
)

const (
	port = 5432
)

func Connect() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		env.GetEnv("DB_NAME"), port, env.GetEnv("POSTGRES_USER"), env.GetEnv("POSTGRES_PASSWORD"), env.GetEnv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&articles.Article{})
	db.AutoMigrate(&users.User{})

	fmt.Println("Successfully connected!")
	return db
}
