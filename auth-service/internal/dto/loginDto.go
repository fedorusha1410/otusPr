package dto

type Login struct {
	Username string `json:"Name"`
	Password string `json:"password"`
}

type SignUp struct {
	Username string `json:"Name"`
	Password string `json:"password"`
}