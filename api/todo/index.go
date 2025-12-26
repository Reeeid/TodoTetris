package todo

import (
	"net/http"

	mdw "github.com/Reeeid/TodoTetris/middleware"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	mdw.AuthJWT(TodoHandler)(w, r)

}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
	case http.MethodGet:
	case http.MethodPut:
	case http.MethodDelete:
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
