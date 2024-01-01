package main

import (
	"echo_blogs/configs"
	"echo_blogs/models"
	"echo_blogs/routes"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := configs.SetupDB()

	_ = db.AutoMigrate(&models.Category{}, &models.Blog{}, &models.User{})

	e := echo.New()

	//e.Use(middleware.Logger())

	e.Static("/public", "public")

	routes.Web(e, db)
	routes.Admin(e, db)
	routes.Api(e, db)

	fmt.Println("Starting Server on: http://localhost:8080")

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
