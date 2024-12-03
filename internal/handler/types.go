package handler

type (
	AuthData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	CreateData struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	Token struct {
		Token     string `json:"token"`
		ExpiresAt int64  `json:"expires_at"`
	}

	Tokens struct {
		AccessToken  Token `json:"access_token"`
		RefreshToken Token `json:"refresh_token"`
	}

	Refresh struct {
		Tokens        Tokens
		IsIpDifferent bool `json:"different_ip"`
	}

	User struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Ip       string `json:"ip"`
		Email    string `json:"email"`
	}
)
