package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/srjchsv/auth-service/internal/app/handlers"
	"github.com/srjchsv/auth-service/internal/pkg/appmetrics"
	"github.com/srjchsv/auth-service/internal/pkg/database"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	// Enable CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// Init metrics
	appmetrics.InitPrometheus(r)

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
