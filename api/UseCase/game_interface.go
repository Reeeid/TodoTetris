package usecase

import "github.com/Reeeid/TodoTetris/Domain/model"

type GameRepository interface {
	SaveGame(Session *model.Session) error
	GameStatus(userID string) (bool, error)
	LoadGame(userID string) (*model.Session, error)
}
