package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/srjchsv/auth-service/internal/app/models"
)

func InitDB() (*gorm.DB, error) {
	dbUser := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbHost := os.Getenv("POSTGRES_HOST")

	dbURI := fmt.Sprintf("host=%v port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost,
		dbPort, dbUser, dbName, dbPassword)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return db, err
	}

	db.AutoMigrate(&models.User{})

	return db, err
}
