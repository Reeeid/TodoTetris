package dto

import "github.com/Reeeid/TodoTetris/pkg/Domain/model"

type TetrisResponse struct {
	IsPlayed bool                     `json:"is_played"`
	Session  *GameSessionLoadResponse `json:"session,omitempty"`
}

type GameSessionSaveRequest struct {
	BoardState string `json:"Board_state"`
	Score      int    `json:"Score"`
}

type GameSessionLoadResponse struct {
	BoardState   string `json:"board_state"`
	Score        int    `json:"Score"`
	LastPlayedAt string `json:"last_played_at"`
}

func (g *GameSessionSaveRequest) ToDomain(username string) *model.Session {
	return &model.Session{
		UserID:     username,
		Score:      g.Score,
		BoardState: g.BoardState,
	}
}
func ToTetrisResponse(IsPlayed bool, m *model.Session) *TetrisResponse {
	var session *GameSessionLoadResponse
	if m != nil {
		session = &GameSessionLoadResponse{
			BoardState:   m.BoardState,
			Score:        m.Score,
			LastPlayedAt: m.LastPlayedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
	}
	return &TetrisResponse{
		IsPlayed: IsPlayed,
		Session:  session,
	}
}
