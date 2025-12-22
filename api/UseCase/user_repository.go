package usecase

type UserRepository interface {
	CreateUser(username string, email string, password string) (string, error)
	GetUserByID(userID string) (map[string]interface{}, error)
}
