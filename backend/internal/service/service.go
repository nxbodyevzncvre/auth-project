package service

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/nxbodyevzncvre/mypackage/internal/config"
	"github.com/nxbodyevzncvre/mypackage/internal/db"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	conf        *config.Config
	authStorage *config.AuthStorage
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		conf:        config.GetConfig(),
		authStorage: &config.AuthStorage{DB: &config.Users{Users: make(map[string]config.User)}},
	}
}

func (h *AuthHandler) PostCard(c *fiber.Ctx) error {
	if db.DB == nil {
		fmt.Printf("DB not initialized")
		return c.Status(fiber.StatusInternalServerError).SendString("DB not initialized")
	}

	card := config.Card{}
	if err := c.BodyParser(&card); err != nil {
		return c.SendString(err.Error())

	}

	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM cards WHERE dish_name = $1)", card.Dish_name).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		c.SendString(err.Error())
	}
	if exists {
		return c.Status(fiber.StatusConflict).SendString("Dish is already created")
	}

	_, err = db.DB.Exec("INSERT INTO cards (dish_name, dish_rating, dish_creator, dish_descr, dish_types) VALUES ($1, $2, $3, $4, $5)", card.Dish_name, card.Dish_rating, card.Dish_creator, card.Dish_descr, card.Dish_types)
	if err != nil {
		return c.SendString(err.Error())

	}
	return c.SendString("Success")

}

func GetAllCards(c *fiber.Ctx) error {
	if db.DB == nil {
		fmt.Printf("DB nor initialized")
		return c.Status(fiber.StatusInternalServerError).SendString("DB not initialized")
	}

	rows, err := db.DB.Query("SELECT id, dish_name, dish_rating, dish_creator, dish_descr, dish_types FROM cards")

	if err != nil {
		return c.SendString(err.Error())
	}
	defer rows.Close()

	var cards []config.Card
	for rows.Next() {
		var card config.Card
		if err := rows.Scan(&card.Id, &card.Dish_name, &card.Dish_rating, &card.Dish_creator, &card.Dish_descr, &card.Dish_types); err != nil {
			return c.SendString("Failed to scan rows")

		}
		cards = append(cards, card)
	}

	return c.JSON(cards)
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	if db.DB == nil {
		fmt.Printf("DB not initialized")
		return c.Status(fiber.StatusInternalServerError).SendString("DB not initialized")
	}
	regReq := config.User{}
	if err := c.BodyParser(&regReq); err != nil {
		return c.SendString(err.Error())
	}

	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", regReq.Email).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return c.SendString(err.Error())
	}

	if exists {
		return c.SendString("E-mail already exists")
	}

	_, err = db.DB.Exec("INSERT INTO users (username, password, email) VALUES ($1, $2, $3)", regReq.Username, regReq.Password, regReq.Email)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.SendString("Success")
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	//getting username and password
	regReq := config.User{}

	if err := c.BodyParser(&regReq); err != nil {
		return c.SendString(err.Error())
	}
	//querry for username existance
	var user config.User
	err := db.DB.QueryRow("SELECT username, password FROM users WHERE username = $1", regReq.Username).Scan(&user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.SendString("User not found")
		}
		return c.SendString(err.Error())
	}

	//if noert exists return error
	//if exists checking password for login
	if regReq.Password != user.Password {
		return c.SendString("Password is incorrect")
	}
	//JWT TOKEN CREATING

	//1. creating data for token (payload chapter)
	payload := jwt.MapClaims{
		"sub": user.Username,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(config.JwtSecretKey)
	if err != nil {
		logrus.WithError(err).Error("JWTT token signing")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(config.LoginResponse{Token: t})
}

func JwtPayloadFromRequest(c *fiber.Ctx) (jwt.MapClaims, bool) {
	jwtToken, ok := c.Context().Value(config.ContextKeyUser).(*jwt.Token)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"jwt_token_context_value": c.Context().Value(config.ContextKeyUser),
		}).Error("wrong type of jwt token in Context")
		return nil, false
	}
	payload, ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok {
		logrus.WithFields(logrus.Fields{
			"jwt_token_claims": jwtToken.Claims,
		}).Error("wrong type of jwt token claims")
		return nil, false
	}
	return payload, true

}

func (h *AuthHandler) Profile(c *fiber.Ctx) error {
	jwtPayload, ok := JwtPayloadFromRequest(c)
	if !ok {
		return c.SendStatus(fiber.StatusUnauthorized)

	}
	rows, err := db.DB.Query("SELECT username FROM users WHERE username = $1", jwtPayload["sub"])
	if err != nil {
		return c.SendString("User not found")
	}

	defer rows.Close()
	var username string
	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return c.JSON(config.ProfileResponse{
		Username: username,
	})
}
