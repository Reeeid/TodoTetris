package entity

import (
	"time"

	"github.com/Reeeid/TodoTetris/api/Domain/model"
)

// gorm models
type Todo struct {
	ID          int64  `gorm:"primaryKey;autoIncrement"`
	UserID      string `gorm:"type:varchar(255);index"`
	User        User   `gorm:"foreignKey:UserID;references:Username;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Subject     string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:text"`
	UUID        string `gorm:"type:uuid"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (e *Todo) ToDomain() *model.Todo {
	return &model.Todo{
		ID:          e.ID,
		UserID:      e.UserID,
		Subject:     e.Subject,
		Description: e.Description,
		UUID:        e.UUID,
	}
}
func FromTodoDomain(m *model.Todo) *Todo {
	return &Todo{
		ID:          m.ID,
		UserID:      m.UserID,
		Subject:     m.Subject,
		Description: m.Description,
		UUID:        m.UUID,
	}
}
