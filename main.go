package main

import (
	"fmt"
	"project/config"
	"project/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	config.ConnectDB()

	routes.Route(app)
	if err := godotenv.Load(); err != nil {
		fmt.Println("error while loading env")
	}
	err := app.Listen(":8080")
	if err != nil {
		fmt.Println("error while listening the port")
	}
	defer config.DB.Close()

}
