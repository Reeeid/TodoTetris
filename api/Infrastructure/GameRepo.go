package infrastructure

import "gorm.io/gorm"

type GameRepoImpl struct {
	db *gorm.DB
}

func NewGameRepo(db *gorm.DB) *GameRepoImpl {
	return &GameRepoImpl{db: db}
}

func (g *GameRepoImpl) SaveGame(UserID string, score int, level int) (string, error) {
	return "nil", nil
}
func (g *GameRepoImpl) GetGameByID(UserID string, score int, level int) (map[string]interface{}, error)

/*SaveGame(userID string, score int, level int) (string, error)
GetGameByID(gameID string) (map[string]interface{}, error)*/
