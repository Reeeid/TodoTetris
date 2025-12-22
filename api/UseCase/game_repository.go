package usecase

type GameRepository interface {
	SaveGame(userID string, score int, level int) (string, error)
	GetGameByID(gameID string) (map[string]interface{}, error)
}
