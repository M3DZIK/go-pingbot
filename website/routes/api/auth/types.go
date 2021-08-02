package auth

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type json map[string]interface{}
