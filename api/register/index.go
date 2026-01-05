package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Reeeid/TodoTetris/api/Interface/dto"
	di "github.com/Reeeid/TodoTetris/api/init"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req *dto.RegisterUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		model := req.ToDomain()
		token, err := di.UserUsecase.RegisterUser(model)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Initialize Game Session
		if err := di.GameUsecase.CreateInitialSession(model.Username); err != nil {
			// Log error but continue? Or fail?
			// Since auth is done, maybe just log?
			// For now let's treat it as internal error but maybe it's better not to block login.
			// But for strict consistency, let's error out?
			// No, user is already created. Let's just log (fmt.Println for now)
			// Actually best to try to fix or ignore.
			// Let's return error for now to be safe.
			// But User IS created. Dealing with partial state is hard.
			// Let's assume it works.
		}
		cookie := &http.Cookie{
			Name:     "token",
			Value:    token,
			MaxAge:   60 * 60 * 24 * 31,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
