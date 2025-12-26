package usecase

//go:generate mockgen -source=$GOFILE -package=mock -destination=../mock/user_repo_mock.go

import "github.com/Reeeid/TodoTetris/Domain/model"

type UserRepository interface {
	CreateUser(m *model.User) error
	FindByUserID(userID string) (bool, *model.User, error)
}
