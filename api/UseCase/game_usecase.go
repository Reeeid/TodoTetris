package usecase

import (
	"time"

	"github.com/Reeeid/TodoTetris/api/Domain/model"
)

type GameUseCase struct {
	repo GameRepository
}

func NewGameUseCase(repo GameRepository) *GameUseCase {
	return &GameUseCase{repo: repo}
}

func (g *GameUseCase) SaveSession(m *model.Session) error {
	m.LastPlayedAt = time.Now()
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

func (g *GameUseCase) CreateInitialSession(userID string) error {
	// Create a session with LastPlayedAt set to yesterday
	// This ensures the user is treated as "Not Played Today" (Morning Routine)
	// but avoids "record not found" errors.
	initialSession := &model.Session{
		UserID:       userID,
		Score:        0,
		BoardState:   "",
		LastPlayedAt: time.Now().AddDate(0, 0, -1),
	}
	return g.repo.SaveGame(initialSession)
}
