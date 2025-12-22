package usecase

type TodoRepository interface {
	CreateTodo(userID string, title string, description string) (string, error)
	GetTodoByID(todoID string) (map[string]interface{}, error)
	UpdateTodo(todoID string, title string, description string) error
	DeleteTodo(todoID string) error
}
