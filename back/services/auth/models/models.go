package models

type UserInfo struct {
	EmailHash    string
	PasswordHash string
}

type RegisterRequest struct {
	EmailHash    string
	PasswordHash string
}

type RegisterResponse struct {
	Token string
}
