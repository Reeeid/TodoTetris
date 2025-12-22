package register

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// 登録処理をここに実装
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
