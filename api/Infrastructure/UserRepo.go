package infrastructure

import "gorm.io/gorm"

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepoImpl {
	return &UserRepoImpl{db: db}
}

func (u *UserRepoImpl) CreateUser(username string, password string) (string, error) {

	return "nil", nil
}

func (u *UserRepoImpl) GetUserByID(userID string, password string) error {
	return nil
}
