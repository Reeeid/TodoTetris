package login

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Reeeid/TodoTetris/Domain/model"
	usecase "github.com/Reeeid/TodoTetris/UseCase"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var loginUser model.LoginUser
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	token, err := usecase.LoginUser(&loginUser)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	//認証処理をここに実装
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   int(time.Hour.Seconds() * 24 * 31),
		HttpOnly: true,
		Secure:   r.TLS != nil,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}
