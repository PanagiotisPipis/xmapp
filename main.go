package main

import (
	"companies-service/handlers"
	"companies-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("companies.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Company{})

	r := gin.Default()

	// Public routes
	r.POST("/auth", func(c *gin.Context) {
		token, _ := handlers.GenerateToken("user")
		c.JSON(200, gin.H{"token": token})
	})

	// Protected routes
	auth := r.Group("/")
	auth.Use(handlers.AuthMiddleware())
	{
		auth.POST("/companies", handlers.CreateCompany(db))
		auth.PATCH("/companies/:id", handlers.PatchCompany(db))
		auth.DELETE("/companies/:id", handlers.DeleteCompany(db))
	}

	// Public read route
	r.GET("/companies/:id", handlers.GetCompany(db))

	r.Run(":8080") // Start server
}
