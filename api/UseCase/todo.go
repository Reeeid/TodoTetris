package usecase

import "net/http"

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// Todo作成ロジックをここに実装
}
func GetTodos(w http.ResponseWriter, r *http.Request) {
	// Todo取得ロジックをここに実装
}
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// Todo更新ロジックをここに実装
}
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// Todo削除ロジックをここに実装
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	// セッション取得ロジックをここに実装
	//スコアも含める
}
