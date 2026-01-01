package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Reeeid/TodoTetris/api/Domain/model"
	"github.com/Reeeid/TodoTetris/api/Interface/dto"
	mdw "github.com/Reeeid/TodoTetris/api/Middleware"
	di "github.com/Reeeid/TodoTetris/api/init"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	mdw.AuthJWT(TodoHandler)(w, r)

}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	//context伝搬でユーザー名をあらかじめ代入しておく
	// アクセス制御のため
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

		var req dto.CreateTodoRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		uuidObj := di.UUIDUsecase.GetTodaysUUID()
		model := req.ToDomain(username, uuidObj.UUID)
		err := di.TodoUsecase.CreateTodo(model)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

	case http.MethodGet:
		model := &model.Todo{
			UserID: username,
		}
		todos, err := di.TodoUsecase.ReadTodos(model)
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
		todo, err := di.TodoUsecase.UpdateTodo(model)
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
			err := di.TodoUsecase.DeleteTodo(&model)
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
