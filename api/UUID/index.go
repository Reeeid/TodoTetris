package uuid

import (
	"encoding/json"
	"net/http"

	"github.com/Reeeid/TodoTetris/Interface/dto"
	di "github.com/Reeeid/TodoTetris/init"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		model := di.UUIDUsecase.GetTodaysUUID()
		res := dto.FromUUIDDomain(model)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
