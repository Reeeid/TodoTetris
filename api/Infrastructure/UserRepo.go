package infrastructure

import (
	"errors"

	"github.com/Reeeid/TodoTetris/Domain/model"
	"github.com/Reeeid/TodoTetris/Infrastructure/entity"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepoImpl {
	return &UserRepoImpl{db: db}
}

func (u *UserRepoImpl) CreateUser(m *model.User) error {
	e := entity.FromUserDomain(m)
	return u.db.Create(&e).Error
}
func (u *UserRepoImpl) FindByUserID(userID string) (bool, *model.User, error) {
	var e entity.User
	err := u.db.Where("username = ?", userID).First(&e).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, nil
		}
		return false, nil, err
	}
	return true, e.ToDomain(), nil
}
