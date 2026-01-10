package dto

import "github.com/Reeeid/TodoTetris/pkg/Domain/model"

type CreateTodoRequest struct {
	Subject     string `json:"subject"`
	Description string `json:"description"`
	UUID        string `json:"uuid"`
}

type CreateTodoBatchRequest struct {
	Todos []CreateTodoRequest `json:"todos"`
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
	UUID        string `json:"uuid"`
}

type UpdateTodoRequest struct {
	ID          int64  `json:"id"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
	UUID        string `json:"uuid"`
}

func (req *UpdateTodoRequest) ToDomain(username string) *model.Todo {
	return &model.Todo{
		ID:          req.ID,
		UserID:      username,
		Subject:     req.Subject,
		Description: req.Description,
		UUID:        req.UUID,
	}
}

//繧｢繝・E繝・E繝・ODO縺ｯ蟾ｮ蛻・蠢懃畑縺ｫTODORESPONSE繧定ｿ斐☆
func ToTodoResponse(m *model.Todo) TodoResponse {
	return TodoResponse{
		ID:          m.ID,
		Subject:     m.Subject,
		Description: m.Description,
		UUID:        m.UUID,
	}
}

//Get Todo縺ｮ繝ｬ繧ｹ繝昴Φ繧ｹ縲繝ｦ繝ｼ繧ｶ繝ｼ蜷阪ｒ繝溘ラ繝ｫ繧ｦ繧ｧ繧｢縺九ｉ縺ｨ繧・

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
			UUID:        m.UUID,
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
