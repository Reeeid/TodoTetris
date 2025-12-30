package dto

import "github.com/Reeeid/TodoTetris/Domain/model"

type TetrisResponse struct {
	IsPlayed bool                     `json:"is_played"`
	Session  *GameSessionLoadResponse `json:"session,omitempty"`
}

type GameSessionSaveRequest struct {
	BoardState string `json:"Board_state"`
	Score      int    `json:"Score"`
}

type GameSessionLoadResponse struct {
	BoardState string `json:"board_state"`
	Score      int    `json:"Score"`
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
			BoardState: m.BoardState,
			Score:      m.Score,
		}
	}
	return &TetrisResponse{
		IsPlayed: IsPlayed,
		Session:  session,
	}
}
