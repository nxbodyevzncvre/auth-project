package config

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "users"
)

type (
	User struct {
		Username string `json:"username"`
		Password string `json:"password"`
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
)

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Host:     host,
			Port:     port,
			User:     user,
			Password: password,
			Dbname:   dbname,
		},
	}
}
