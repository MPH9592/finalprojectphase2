package main

import (
	"log"
	"rental-system/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connect to the database
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize the handler
	handler.InitializeHandler(db)

	// Create a Gin router
	r := gin.Default()

	// User routes
	r.POST("/users", handler.CreateUser)
	r.GET("/users", handler.GetUsers)
	r.GET("/users/:id", handler.GetUser)

	// Book routes
	r.POST("/books", handler.CreateBook)
	r.GET("/books", handler.GetBooks)
	r.GET("/books/:id", handler.GetBook)

	// Rental routes
	r.POST("/rentals", handler.CreateRental)
	r.GET("/rentals", handler.GetRentals)

	// Run the server
	r.Run(":8080")
}
