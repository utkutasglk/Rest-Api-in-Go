package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/utkutasglk/Rest-Api-in-Go/database"
	"github.com/utkutasglk/Rest-Api-in-Go/routes"
)

func welcome(c *fiber.Ctx) error{
	return c.SendString("Welcome to Api")
}

func setupRoutes(app *fiber.App) {
	// welcome
	app.Get("/api",welcome)

	app.Post("/api/users", routes.CreateUser)
}

func main(){
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}