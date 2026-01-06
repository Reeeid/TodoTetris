package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Reeeid/TodoTetris/pkg/Domain/model"
	"github.com/Reeeid/TodoTetris/pkg/Interface/dto"
	mdw "github.com/Reeeid/TodoTetris/pkg/Middleware"
	di "github.com/Reeeid/TodoTetris/pkg/init"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	mdw.AuthJWT(TodoHandler)(w, r)

}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	val := r.Context().Value(mdw.UserKey)
	if val == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	username, ok := val.(string)
	if !ok {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodPost:

		var batchReq dto.CreateTodoBatchRequest
		if err := json.NewDecoder(r.Body).Decode(&batchReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		responseTodos := make([]dto.TodoResponse, 0)

		for _, reqItem := range batchReq.Todos {
			uuidObj := di.GetUUIDUsecase().GetTodaysUUID()
			model := reqItem.ToDomain(username, uuidObj.UUID)
			err := di.GetTodoUsecase().CreateTodo(model)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			responseTodos = append(responseTodos, dto.ToTodoResponse(model))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"todos": responseTodos}); err != nil {
		}

	case http.MethodGet:
		model := &model.Todo{
			UserID: username,
		}
		todos, err := di.GetTodoUsecase().ReadTodos(model)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := dto.ToReadTodoResponse(todos)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	case http.MethodPut:
		var req dto.UpdateTodoRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		model := req.ToDomain(username)
		todo, err := di.GetTodoUsecase().UpdateTodo(model)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := dto.ToTodoResponse(todo)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return

	case http.MethodDelete:
		var req dto.DeleteTodoRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		models := req.ToDomain(username)
		for _, model := range models {
			err := di.GetTodoUsecase().DeleteTodo(&model)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
