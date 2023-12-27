package main

import (
	"github.com/baddar99/go-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
    app := fiber.New()

	setupRoutes(app)



    app.Listen(":3000")
}