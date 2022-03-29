package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ittechman101/go-pos/database"
	"github.com/ittechman101/go-pos/pos"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()
	//	defer database.DB.Close()

	//	api := app.Group("/api")
	pos.Register(app, database.DB)

	log.Fatal(app.Listen(":5000"))
}
