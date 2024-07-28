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
		Username string `json: "username"`
		Password string `json: "password"`
	}

	Users struct {
		Users map[string]User
	}
	DBConfig struct {
		host     string
		port     int
		user     string
		password string
		dbname   string
	}
	Config struct {
		DB *DBConfig
	}
)

func getConfig() *Config {
	return &Config{
		DB: &DBConfig{
			host:     host,
			port:     port,
			user:     user,
			password: password,
			dbname:   dbname,
		},
	}
}
