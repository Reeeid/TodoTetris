package entity

import (
	"time"

	"github.com/Reeeid/TodoTetris/Domain/model"
)

type User struct {
	Username     string `gorm:"primaryKey;type:varchar(255)"`
	PasswordHash string `gorm:"type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) ToDomain() *model.User {
	return &model.User{
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
	}
}

func FromUserDomain(m *model.User) *User {
	return &User{
		Username:     m.Username,
		PasswordHash: m.PasswordHash,
	}
}
