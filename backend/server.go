package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
)

func main() {

	if err := connectDB(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Post("/sign-up", func(c *fiber.Ctx) error {
		u := new(User)

		if err := c.BodyParser(u); err != nil {
			return c.Status(404).SendString(err.Error())
		}

		output, err := db.Query("INSERT INTO users (username, password) VALUES ($1, $2)", u.Username, u.Password)

		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		log.Println(output)

		return c.JSON((c.BodyParser(u)))

	})

	app.Get("/get_users", func(c *fiber.Ctx) error {
		u := new(User)
		if err := c.BodyParser(u); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		data, err := db.Query("SELECT * FROM users")
		if err != nil {
			return c.Status(400).SendString(err.Error())

		}
		defer data.Close()

		res := Users{}

		for data.Next() {
			user := User{}
			if err := data.Scan(&user.Username, &user.Password, &user.Id); err != nil {
				return c.Status(400).SendString(err.Error())
			}
			res.Users = append(res.Users, user)
		}
		return c.JSON(res)
	})

	app.Post("sign-in/:username", func(c *fiber.Ctx) error {
		u := new(User)
		if err := c.BodyParser(u); err != nil {
			return c.Status(404).SendString(err.Error())
		}

		data, err := db.Query("SELECT * FROM users WHERE (username, password) = $1, $2", u.Username, u.Password)
		if err != nil {
			return c.Status(404).SendString(err.Error())
		}

		defer data.Close()

		res := Users{}

		for data.Next() {
			user := User{}
			if err := data.Scan(&user.Username, &user.Password, &user.Id); err != nil {
				return c.Status(404).SendString(err.Error())
			}
			res.Users = append(res.Users, user)
		}

		return c.JSON(res)
	})

	app.Listen(":8080")

}
