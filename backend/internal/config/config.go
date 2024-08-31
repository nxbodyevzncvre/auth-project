package config

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

var JwtSecretKey = []byte("pd[asfckjiogiotijt]")
var ContextKeyUser = "user"

type (
	User struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email 	 string `json:""`		
	}

	Card struct {
		Dish_name    string `bson:"dish_name"`
		Dish_rating  int    `bson:"dish_rating"`
		Dish_creator string `bson:"dish_creator"`
		Dish_descr   string `bson:"dish_descr"`
		Dish_types   string `bson:"dish_types"`
		ImageID      primitive.ObjectID `bson:"image_id"`
	}

	Svg struct {
		Url string `json:"url"`
	}

	AuthStorage struct {
		DB *Users
	}

	Users struct {
		Users map[string]User
	}

	DBConfig struct {
		Host     string
		Port     int
		User     string
		Password string
		Dbname   string
	}
	ProfileResponse struct {
		Username string `json:"username"`
	}

	Config struct {
		DB *DBConfig
	}

	LoginResponse struct {
		Token string `json:"access_token"`
	}
)

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "admin",
			Dbname:   "postgres",
		},
	}
}
