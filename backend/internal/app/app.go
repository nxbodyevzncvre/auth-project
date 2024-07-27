package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	if err := connectDB(); err != nil {
		log.Fatal(err)

	}
	app := fiber.New()

	app.Listen(":8080")
}
