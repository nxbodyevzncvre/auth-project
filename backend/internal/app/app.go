package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	if err := database.connectDBase(); err != nil {
		log.Fatal(err)

	}
	app := fiber.New()

	app.Listen(":8080")
}
