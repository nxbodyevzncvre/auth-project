package main

import (
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nxbodyevzncvre/mypackage/internal/config"
	"github.com/nxbodyevzncvre/mypackage/internal/db"
	"github.com/nxbodyevzncvre/mypackage/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := db.Init(); err != nil {
		logrus.Fatal(err)
		fmt.Printf("failed init db %s", err)
	}

	app := fiber.New()

	app.Use(cors.New())

	authHandlers := service.NewAuthHandler()

	app.Post("/register", authHandlers.Register)
	app.Post("/login", authHandlers.Login)
	app.Post("/create-card", authHandlers.PostCard)
	app.Get("/cards", service.GetAllCards)
	authtorizedGroup := app.Group("")
	authtorizedGroup.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: config.JwtSecretKey,
		},
		ContextKey: config.ContextKeyUser,
	}))

	authtorizedGroup.Get("/profile", authHandlers.Profile)

	app.Listen(":8080")
}
