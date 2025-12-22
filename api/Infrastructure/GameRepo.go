package infrastructure

type GameRepo struct {
	db DBprovider
}

func NewGameRepo(db DBprovider) *GameRepo {
	return &GameRepo{db: db}
}

func (g *GameRepo) SaveGame(UserID string, score int, level int) (string, error) {
	return "nil", nil
}
func (g *GameRepo) GetGameByID(UserID string, score int, level int) (map[string]interface{}, error)

/*SaveGame(userID string, score int, level int) (string, error)
GetGameByID(gameID string) (map[string]interface{}, error)*/
