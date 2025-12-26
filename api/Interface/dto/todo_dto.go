package dto

import "github.com/Reeeid/TodoTetris/Domain/model"

//Todo作成と返答

type CreateTodoRequest struct {
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

func (req *CreateTodoRequest) ToDomain(uuid string) *model.Todo {
	return &model.Todo{
		Subject:     req.Subject,
		Description: req.Description,
		UUID:        uuid,
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

func (req *UpdateTodoRequest) ToDomain() *model.Todo {
	return &model.Todo{
		ID:          req.ID,
		Subject:     req.Subject,
		Description: req.Description,
	}
}

func ToTodoResponse(m *model.Todo) TodoResponse {
	return TodoResponse{
		ID:          m.ID,
		Subject:     m.Subject,
		Description: m.Description,
	}
}

type ReadTodoRequest struct {
	IDs []int64 `json:"ids"`
}

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
