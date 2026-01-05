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
		// Read body to allow potential double-decode or just try Batch format logic
		// We'll support the new Batch format: { "todos": [ ... ] }
		// If the frontend sends this, we process it.
		// NOTE: To support legacy single item (if any), we could verify.
		// But let's assume we move forward with Batch as primary or just support Batch struct.

		var batchReq dto.CreateTodoBatchRequest
		if err := json.NewDecoder(r.Body).Decode(&batchReq); err != nil {
			// If decode fails, maybe it was single? Or bad JSON.
			// Given user instructions, let's assume we fix frontend same time.
			// But 'Decode' consumes reader.
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// If 'todos' is empty, maybe it was a single request { "subject": ... } which resulted in empty 'todos' slice?
		// Let's handle just the batch for now as requested "Change API".

		responseTodos := make([]dto.TodoResponse, 0)

		// Process loop
		for _, reqItem := range batchReq.Todos {
			uuidObj := di.UUIDUsecase.GetTodaysUUID()
			model := reqItem.ToDomain(username, uuidObj.UUID)
			err := di.TodoUsecase.CreateTodo(model)
			if err != nil {
				// On error, we stop? Or continue?
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			responseTodos = append(responseTodos, dto.ToTodoResponse(model))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"todos": responseTodos}); err != nil {
			// log error
		}

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
