package infrastructure

import (
	"time"

	"github.com/Reeeid/TodoTetris/api/Domain/model"
	"github.com/Reeeid/TodoTetris/api/Infrastructure/entity"
	"gorm.io/gorm"
)

type GameRepoImpl struct {
	db *gorm.DB
}

func NewGameRepo(db *gorm.DB) *GameRepoImpl {
	return &GameRepoImpl{db: db}
}

func (g *GameRepoImpl) SaveGame(m *model.Session) error {
	e := entity.FromSessionDomain(m)
	if err := g.db.Save(e).Error; err != nil {
		return err
	}
	return nil
}

func (g *GameRepoImpl) GameStatus(userID string) (bool, error) {
	var e entity.GameSession
	err := g.db.Select("last_played_at").Where("user_id = ?", userID).First(&e).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := time.Now().In(jst)
	lastJST := e.LastPlayedAt.In(jst)
	isPlayedToday := nowJST.Year() == lastJST.Year() && nowJST.Month() == lastJST.Month() && nowJST.Day() == lastJST.Day()

	return isPlayedToday, nil
}

func (g *GameRepoImpl) LoadGame(userID string) (*model.Session, error) {
	var e entity.GameSession
	err := g.db.Where("user_id = ?", userID).First(&e).Error
	if err != nil {
		return nil, err
	}
	return e.ToDomain(), nil
}
