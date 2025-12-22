package Todo

import (
	"net/http"

	mdw "github.com/Reeeid/TodoTetris/middleware"
	usecase "github.com/Reeeid/TodoTetris/usecase"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	mdw.AuthJWT(TodoHandler)(w, r)

}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		usecase.CreateTodo(w, r)
	case http.MethodGet:
		usecase.GetTodos(w, r)
	case http.MethodPut:
		usecase.UpdateTodo(w, r)
	case http.MethodDelete:
		usecase.DeleteTodo(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
