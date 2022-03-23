package database

import (
	"github.com/davisbento/gorm-api/core/articles"
	"github.com/davisbento/gorm-api/core/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"
)

const (
	host     = "database"
	port     = 5432
	user     = "davisbento"
	password = "davisbento_pass"
	dbname   = "articles"
)

func Connect() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
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
