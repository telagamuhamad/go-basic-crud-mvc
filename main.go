package main

import (
	"fmt"
	"log"
	"testing-app/migrations"
	"testing-app/models"
	"testing-app/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/testing_go?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	models.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connected!")
	models.DB.AutoMigrate(&models.User{})
}

func main() {
	initDatabase()

	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// migrations
	migrations.Migrate(models.DB)

	log.Fatal(r.Run(":8080"))
}
