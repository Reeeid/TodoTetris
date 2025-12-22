package infrastructure

type UserRepo struct {
	db DBprovider
}

func NewUserRepo(db DBprovider) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) CreateUser(username string, password string) (string, error) {
	return "nil", nil
}

func (u *UserRepo) GetUserByID(userID string, password string) error {
	return nil
}
