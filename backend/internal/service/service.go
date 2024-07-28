package service

import (
	"fmt"
	"database/sql"
	"github.com/gofiber/fiber/v2" 
	"github.com/nxbodyevzncvre/mypackage/internal/config"
	"github.com/nxbodyevzncvre/mypackage/internal/db"
    jwt "github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"time"
)

type AuthHandler struct{
	conf *config.Config
	authStorage *config.AuthStorage
}
	
func NewAuthHandler() *AuthHandler{
	return &AuthHandler{
		conf: config.GetConfig(),
		authStorage: &config.AuthStorage{DB: &config.Users{Users: make(map[string]config.User)}},
	}
}
var jwtSecretKey = []byte("foapfaophjfdm")

func (h *AuthHandler) Register(c *fiber.Ctx) error{
	if db.DB == nil{
		fmt.Printf("DB not initialized")
		return c.Status(fiber.StatusInternalServerError).SendString("DB not initialized")
	}
	regReq := config.User{}
	if err := c.BodyParser(&regReq); err != nil{
		return c.SendString(err.Error())
	}

	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", regReq.Username).Scan(&exists)
	if err !=nil && err != sql.ErrNoRows{
		return c.SendString(err.Error())
	}

	if exists{
		return c.Status(fiber.StatusConflict).SendString("User already exists")
	}

	_, err = db.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", regReq.Username, regReq.Password)
	if err != nil{
		return c.SendString(err.Error())
	}

	return c.SendString("Success")
}

func (h *AuthHandler) Login(c *fiber.Ctx) error{
	//getting username and password
	regReq := config.User{}

	if err := c.BodyParser(&regReq); err != nil{
		return c.SendString(err.Error())
	}
	//querry for username existance
	var user config.User
	err := db.DB.QueryRow("SELECT username, password FROM users WHERE username = $1", regReq.Username).Scan(&user.Username, &user.Password)
	if err != nil{
		if err == sql.ErrNoRows{
			return c.SendString("User not found")
		}
		return c.SendString(err.Error())
	}

	
	//if noert exists return error
	//if exists checking password for login
	if user.Password != regReq.Password {
		return c.SendString("Password is inccorect")
	}
	//JWT TOKEN CREATING

	//1. creating data for token (payload chapter)
	payload := jwt.MapClaims{
		"sub": user.Username,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	
	
	t, err := token.SignedString(jwtSecretKey)
	if err != nil{
		logrus.WithError(err).Error("JWTT token signing")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(config.LoginResponse{AccessToken: t})
}