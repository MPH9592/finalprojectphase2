package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Response for handling errors
type ErrorResponse struct {
	Message string `json:"message"`
}

// Database instance
var db *gorm.DB

// InitializeHandler initializes the handler with a database instance
func InitializeHandler(database *gorm.DB) {
	db = database
}

// CreateUser handles user creation
func CreateUser(c *gin.Context) {
	type UserRequest struct {
		Email    string  `json:"email" binding:"required,email"`
		Password string  `json:"password" binding:"required"`
	}

	var userRequest UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	user := User{
		Email:         userRequest.Email,
		Password:      userRequest.Password,
		DepositAmount: 0,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUsers handles fetching all users
func GetUsers(c *gin.Context) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUser handles fetching a specific user by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateBook handles book creation
func CreateBook(c *gin.Context) {
	type BookRequest struct {
		Name              string  `json:"name" binding:"required"`
		StockAvailability int     `json:"stock_availability" binding:"required"`
		RentalCosts       float64 `json:"rental_costs" binding:"required"`
		Category          string  `json:"category" binding:"required"`
	}

	var bookRequest BookRequest
	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	book := Book{
		Name:              bookRequest.Name,
		StockAvailability: bookRequest.StockAvailability,
		RentalCosts:       bookRequest.RentalCosts,
		Category:          bookRequest.Category,
	}

	if err := db.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// GetBooks handles fetching all books
func GetBooks(c *gin.Context) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to fetch books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// GetBook handles fetching a specific book by ID
func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// CreateRental handles the creation of a rental record
func CreateRental(c *gin.Context) {
	type RentalRequest struct {
		UserID uint `json:"user_id" binding:"required"`
		BookID uint `json:"book_id" binding:"required"`
	}

	var rentalRequest RentalRequest
	if err := c.ShouldBindJSON(&rentalRequest); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	rental := RentalHistory{
		UserID:   rentalRequest.UserID,
		BookID:   rentalRequest.BookID,
		RentedOn: "",
	}

	if err := db.Create(&rental).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create rental"})
		return
	}

	c.JSON(http.StatusCreated, rental)
}

// GetRentals handles fetching rental history
func GetRentals(c *gin.Context) {
	var rentals []RentalHistory
	if err := db.Preload("User").Preload("Book").Find(&rentals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to fetch rentals"})
		return
	}
	c.JSON(http.StatusOK, rentals)
}
