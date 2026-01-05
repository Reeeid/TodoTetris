package usecase

import "github.com/Reeeid/TodoTetris/pkg/Domain/model"

type TodoRepository interface {
	CreateTodo(todo *model.Todo) error
	ReadTodo(UserID string) ([]model.Todo, error)
	UpdateTodo(todo *model.Todo) (*model.Todo, error)
	DeleteTodo(todo *model.Todo) error
}
