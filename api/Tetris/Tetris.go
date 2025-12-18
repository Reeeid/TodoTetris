package Tetris

import (
	"net/http"

	mdw "github.com/Reeeid/TodoTetris/middleware"
	svc "github.com/Reeeid/TodoTetris/service"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	mdw.AuthJWT(nil)(w, r)
}

func TetrisHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		svc.GetSession(w, r)
	case http.MethodDelete:
		svc.DeleteTodo(w, r)
	}
}
