package config

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

var jwtSecretKey = []byte("pd[asfckjiogiotijt]")

type (
	User struct {
		Username string `json:"username"`
		Password string `json:"password"`
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

	Config struct {
		DB *DBConfig
	}

	LoginResponse struct{
		AccessToken string `json:"access_token"`
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


