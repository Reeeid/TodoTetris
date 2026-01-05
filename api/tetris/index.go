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
		//スチEEタスチェチEでセチEョンを返すか判断する
		status, err := di.GameUsecase.GameStatus(model)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// プレイ済みかどうかにかかわらず、前回のセッション情報を取得する（ペナルティ計算用）
		session, err := di.GameUsecase.LoadGame(model)
		// エラーは無視（初回ユーザーなど）

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
		if err := di.GameUsecase.SaveSession(model); err != nil {
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
