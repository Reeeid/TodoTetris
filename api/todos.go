package api

import (
	"net/http"

	mdw "github.com/Reeeid/TodoTetris/middleware"
	svc "github.com/Reeeid/TodoTetris/service"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	mdw.SupabaseJWT(TodoHandler)(w, r)

}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		svc.CreateTodo(w, r)
	case http.MethodGet:
		svc.GetTodos(w, r)
	case http.MethodPut:
		svc.UpdateTodo(w, r)
	case http.MethodDelete:
		svc.DeleteTodo(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
