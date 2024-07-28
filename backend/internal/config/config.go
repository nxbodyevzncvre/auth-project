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
)
