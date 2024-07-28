package backend

import (
	"github.com/gofiber/fiber"
	"github.com/nxbodyevcre/internal/db"
)

func main() {
	app := fiber.New()
	config := config.getConfig()
	db.Init(config)
	app.Listen(":3000")
}
