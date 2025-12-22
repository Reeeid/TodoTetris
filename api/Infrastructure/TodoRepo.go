package infrastructure

type TodoRepo struct {
	db DBprovider
}

func NewTodoRepo(db DBprovider) *TodoRepo {
	return &TodoRepo{db: db}
}

func (t *TodoRepo) CreateTodo(userID string, title string, description string) (string, error) {
	// Implementation here
	return "", nil
}
func (t *TodoRepo) GetTodo(todoID string) (map[string]interface{}, error) {
	// Implementation here
	return nil, nil
}
func (t *TodoRepo) UpdateTodo(todoID, title string, description string) error {
	// Implementation here
	return nil
}
func (t *TodoRepo) DeleteTodo(todoID string) error {
	// Implementation here
	return nil
}
