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
	mdw.AuthJWT(TetrisHandler)(w, r)
}

func TetrisHandler(w http.ResponseWriter, r *http.Request) {

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
	case http.MethodGet:
		model := &model.Session{
			UserID: username,
		}
		status, err := di.GetGameUsecase().GameStatus(model)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 繝励Ξ繧､貂医∩縺九←縺・°縺ｫ縺九°繧上ｉ縺壹∝燕蝗槭・繧ｻ繝・す繝ｧ繝ｳ諠・ｱ繧貞叙蠕励☆繧具ｼ医・繝翫Ν繝・ぅ險育ｮ礼畑・・
		session, err := di.GetGameUsecase().LoadGame(model)

		res := dto.ToTetrisResponse(status, session)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return

	case http.MethodPost:
		var req dto.GameSessionSaveRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		model := req.ToDomain(username)
		if err := di.GetGameUsecase().SaveSession(model); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
