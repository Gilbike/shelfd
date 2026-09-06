package auth

type authenticateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
