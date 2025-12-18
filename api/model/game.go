package model

import (
	"encoding/json"
	"time"
)

type (
	Session struct {
		UserID      string          // ユーザーID
		Score       int             `json:"score"`       // 現在のスコア
		UUID        string          `json:"uuid"`        // 今日の共通UUID
		BoardState  json.RawMessage `json:"board_state"` // 盤面（ここにミノごとのUUIDが含まれる）
		LastUpdated time.Time       `json:"last_updated"`
	}

	ClearCheckRequest struct {
		UUID []string `json:"uuid"` //フロントエンドでUUIDが消えたか判断して消えたUUIDが送られる。
	}
)
