package tetris

import (
	"net/http"

	mdw "github.com/Reeeid/TodoTetris/middleware"
	usecase "github.com/Reeeid/TodoTetris/usecase"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	mdw.AuthJWT(TetrisHandler)(w, r)
}

func TetrisHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		usecase.GetSession(w, r)
	case http.MethodDelete:
		usecase.DeleteTodo(w, r)
	}
}
