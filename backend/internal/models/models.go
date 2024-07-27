package models

type (
	AuthHandler struct {
		storage *Users
	}

	User struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Users struct {
		Users []User `json:"users"`
	}

	RequestData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	ProfileResponse struct {
		Username string `json:"username"`
	}

	LoginResponse struct {
		AccessToken string `json:"access_token"`
	}
)
