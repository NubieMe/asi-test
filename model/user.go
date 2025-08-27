package model

type UserData struct {
	Realname string `json:"realname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterPayload struct {
	Username string `json:"username"`
	Realname string `json:"realname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
