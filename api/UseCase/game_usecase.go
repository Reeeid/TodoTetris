package usecase

import "github.com/Reeeid/TodoTetris/api/Domain/model"

type GameUseCase struct {
	repo GameRepository
}

func NewGameUseCase(repo GameRepository) *GameUseCase {
	return &GameUseCase{repo: repo}
}

func (g *GameUseCase) SaveSession(m *model.Session) error {
	err := g.repo.SaveGame(m)
	if err != nil {
		return err
	}
	return nil
}
func (g *GameUseCase) GameStatus(m *model.Session) (bool, error) {
	ok, err := g.repo.GameStatus(m.UserID)
	if err != nil {
		return false, err
	}
	return ok, nil
}
func (g *GameUseCase) LoadGame(m *model.Session) (*model.Session, error) {
	session, err := g.repo.LoadGame(m.UserID)
	if err != nil {
		return nil, err
	}
	return session, nil
}
