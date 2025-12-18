package model

type (
	TODO struct {
		ID          int64  `json:"id"`
		Subject     string `json:"subject"`
		Description string `json:"description"`
		UserID      string `json:"user_id"`
		UUID        string `json:"uuid"`
	}

	CreateTodoRequest struct {
		Subject     string `json:"subject"`
		Description string `json:"description"`
		UserID      string `json:"user_id"`
	}

	GetTodaysUUIDresponse struct {
		UUID string `json:"uuid"`
	}

	UpdateTodoRequest struct {
		ID          int64  `json:"id"`
		Subject     string `json:"subject"`
		Description string `json:"description"`
		UserID      string `json:"user_id"`
	}
)
