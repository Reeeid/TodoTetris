package dto

import "github.com/Reeeid/TodoTetris/api/Domain/model"

type RegisterUserRequest struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

func (r *RegisterUserRequest) ToDomain() *model.User {
	return &model.User{
		Username:     r.Username,
		PasswordHash: r.PasswordHash,
	}
}

type LoginUserRequest struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

func (l *LoginUserRequest) ToDomain() *model.User {
	return &model.User{
		Username:     l.Username,
		PasswordHash: l.PasswordHash,
	}
}
