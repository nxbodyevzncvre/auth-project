package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nxbodyevzncvre/mypackage/internal/db"
	"github.com/nxbodyevzncvre/mypackage/internal/service"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := db.Init(); err !=nil{
		logrus.Fatal(err)
		fmt.Printf("failed init db %s", err)
	}

	app := fiber.New()
	
	authHandlers := service.NewAuthHandler()
	
	app.Post("/register", authHandlers.Register)
	app.Post("/login", authHandlers.Login)

	app.Listen(":8080")
}
