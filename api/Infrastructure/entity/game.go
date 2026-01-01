package entity

import (
	"time"
	_ "time"

	"github.com/Reeeid/TodoTetris/api/Domain/model"
	_ "github.com/Reeeid/TodoTetris/api/Domain/model"
)

type GameSession struct {
	UserID       string `gorm:"primaryKey;type:varchar(255)"`
	User         User   `gorm:"foreignKey:UserID;references:Username;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Score        int    `gorm:"default:0"`
	BoardState   string `gorm:"type:text"`
	LastPlayedAt time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (e *GameSession) ToDomain() *model.Session {
	return &model.Session{
		UserID:       e.UserID,
		Score:        e.Score,
		BoardState:   e.BoardState,
		LastPlayedAt: e.LastPlayedAt,
	}
}

func FromSessionDomain(m *model.Session) *GameSession {
	return &GameSession{
		UserID:       m.UserID,
		Score:        m.Score,
		BoardState:   m.BoardState,
		LastPlayedAt: m.LastPlayedAt,
	}
}
