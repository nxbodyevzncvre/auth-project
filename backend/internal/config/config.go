package config

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
		Email    string `json:"email"`
	}

	Card struct {
		Id			 int `json:"id"`
		Dish_name    string `json:"dish_name"`
		Dish_rating  int    `json:"dish_rating"`
		Dish_creator string `json:"dish_creator"`
		Dish_descr   string `json:"dish_descr"`
		Dish_types   string `json:"dish_types"`
	}

	Svg struct{
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
	ProfileResponse struct{
		Username string `json:"username"`
	}

	Config struct {
		DB *DBConfig
	}

	LoginResponse struct{
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


