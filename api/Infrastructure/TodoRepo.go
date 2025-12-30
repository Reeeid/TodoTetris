package infrastructure

import (
	"github.com/Reeeid/TodoTetris/Domain/model"
	"github.com/Reeeid/TodoTetris/Infrastructure/entity"
	"gorm.io/gorm"
)

type TodoRepoImpl struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) *TodoRepoImpl {
	return &TodoRepoImpl{db: db}
}

func (t *TodoRepoImpl) CreateTodo(todo *model.Todo) error {

	e := entity.FromTodoDomain(todo)
	if err := t.db.Create(&e).Error; err != nil {
		return err
	}
	return nil
}
func (t *TodoRepoImpl) ReadTodo(userID string) ([]model.Todo, error) {
	var todos []entity.Todo
	result := t.db.Where("user_id = ?", userID).Find(&todos)
	err := result.Error
	if err != nil {
		return nil, err
	}
	var modelTodos []model.Todo
	for _, todo := range todos {
		modelTodos = append(modelTodos, *todo.ToDomain())
	}
	return modelTodos, nil
}
func (t *TodoRepoImpl) UpdateTodo(todo *model.Todo) (*model.Todo, error) {
	userID := todo.UserID
	id := todo.ID
	var e entity.Todo
	if err := t.db.Model(&e).Where("id = ? AND user_id = ?", id, userID).Updates(map[string]string{
		"subject":     todo.Subject,
		"description": todo.Description,
	}).Error; err != nil {
		return nil, err
	}
	//todo返すでもいいかもしれない
	if err := t.db.First(&e, id).Error; err != nil {
		return nil, err
	}
	//
	return e.ToDomain(), nil
}

func (t *TodoRepoImpl) DeleteTodo(todo *model.Todo) error {
	userID := todo.UserID
	uuid := todo.UUID
	if err := t.db.Where("uuid = ? AND user_id = ?", uuid, userID).Delete(&entity.Todo{}).Error; err != nil {
		return err
	}
	return nil
}

/*
	CreateTodo(todo *model.Todo) error
	ReadTodo(todo *model.Todo) ([]model.Todo, error)
	UpdateTodo(todo *model.Todo) (*model.Todo, error)
	DeleteTodo(todo *model.Todo) (*model.Todo, error)
*/
