package request

type AuthUser struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
