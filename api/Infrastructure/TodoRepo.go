package infrastructure

import "gorm.io/gorm"

type TodoRepoImpl struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) *TodoRepoImpl {
	return &TodoRepoImpl{db: db}
}

func (t *TodoRepoImpl) CreateTodo(userID string, title string, description string) (string, error) {
	// Implementation here
	return "", nil
}
func (t *TodoRepoImpl) GetTodo(todoID string) (map[string]interface{}, error) {
	// Implementation here
	return nil, nil
}
func (t *TodoRepoImpl) UpdateTodo(todoID, title string, description string) error {
	// Implementation here
	return nil
}
func (t *TodoRepoImpl) DeleteTodo(todoID string) error {
	// Implementation here
	return nil
}
