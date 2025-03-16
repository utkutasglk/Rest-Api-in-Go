package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/utkutasglk/Rest-Api-in-Go/database"
)

func welcome(c *fiber.Ctx) error{

	return c.SendString("Welcome to Api")
}

func main(){
	database.ConnectDb()

	app := fiber.New()

	app.Get("/api",welcome)

	log.Fatal(app.Listen(":3000"))
}