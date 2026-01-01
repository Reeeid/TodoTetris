package dto

import "github.com/Reeeid/TodoTetris/api/Domain/model"

type CreateTodoRequest struct {
	Subject     string `json:"subject"`
	Description string `json:"description"`
	UUID        string `json:"uuid"`
}

func (req *CreateTodoRequest) ToDomain(username string, UUID string) *model.Todo {
	return &model.Todo{
		Subject:     req.Subject,
		Description: req.Description,
		UUID:        UUID,
		UserID:      username,
	}
}

type TodoResponse struct {
	ID          int64  `json:"id"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

type UpdateTodoRequest struct {
	ID          int64  `json:"id"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

func (req *UpdateTodoRequest) ToDomain(username string) *model.Todo {
	return &model.Todo{
		ID:          req.ID,
		UserID:      username,
		Subject:     req.Subject,
		Description: req.Description,
	}
}

//アチE�EチE�EチEODOは差刁E��応用にTODORESPONSEを返す
func ToTodoResponse(m *model.Todo) TodoResponse {
	return TodoResponse{
		ID:          m.ID,
		Subject:     m.Subject,
		Description: m.Description,
	}
}

//Get Todoのレスポンス　ユーザー名をミドルウェアからとる

type ReadTodoResponse struct {
	Todos []TodoResponse `json:"todos"`
}

func ToReadTodoResponse(models []model.Todo) ReadTodoResponse {
	todos := make([]TodoResponse, len(models))
	for i, m := range models {
		todos[i] = TodoResponse{
			ID:          m.ID,
			Subject:     m.Subject,
			Description: m.Description,
		}
	}
	return ReadTodoResponse{Todos: todos}
}

type DeleteTodoRequest struct {
	UUIDs []string `json:"uuids"`
}

func (d *DeleteTodoRequest) ToDomain(username string) []model.Todo {
	todos := make([]model.Todo, len(d.UUIDs))
	for i, j := range d.UUIDs {
		todos[i] = model.Todo{
			UserID: username,
			UUID:   j,
		}
	}
	return todos
}
