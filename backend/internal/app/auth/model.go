package auth

type user struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
