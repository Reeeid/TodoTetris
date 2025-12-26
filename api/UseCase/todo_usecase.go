package usecase

import "github.com/Reeeid/TodoTetris/Domain/model"

type TodoUseCase struct {
	repo TodoRepository
}

func NewTodoUseCase(repo TodoRepository) *TodoUseCase {
	return &TodoUseCase{repo: repo}
}

func (t *TodoUseCase) CreateTodo(m *model.Todo) error {
	if err := t.repo.CreateTodo(m); err != nil {
		return err
	}
	return nil
}
func (t *TodoUseCase) ReadTodos(m *model.Todo) ([]model.Todo, error) {
	result, err := t.repo.ReadTodo(m.UserID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (t *TodoUseCase) UpdateTodo(m *model.Todo) (*model.Todo, error) {
	result, err := t.repo.UpdateTodo(m)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t *TodoUseCase) DeleteTodo(m *model.Todo) error {
	err := t.repo.DeleteTodo(m)
	if err != nil {
		return err
	}
	return nil
}
