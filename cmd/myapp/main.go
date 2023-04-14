package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/srjchsv/auth-service/internal/app/handlers"
	"github.com/srjchsv/auth-service/internal/app/models"
)

func main() {

	dbUser := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbHost := os.Getenv("POSTGRES_HOST")

	dbURI := fmt.Sprintf("host=%v port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost,
		dbPort, dbUser, dbName, dbPassword)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&models.User{})

	r := gin.Default()
	// Enable CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.POST("/register", handlers.Register(db))
	r.POST("/login", handlers.Login(db))

	users := r.Group("/")
	users.Use(handlers.AuthMiddleware())
	{
		users.GET("/user", handlers.GetUser(db))
		users.PUT("/user", handlers.UpdateUser(db))
		users.DELETE("/user", handlers.DeleteUser(db))
	}
	r.GET("/users", handlers.ListUsers(db))

	r.Run()
}
