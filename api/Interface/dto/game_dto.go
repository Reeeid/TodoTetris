package dto

import "github.com/Reeeid/TodoTetris/Domain/model"

type GameStatusResponse struct {
	IsPlayed bool `json:"isplayed"`
}

type GameSessionSaveRequest struct {
	BoardState string `json:"Board_state"`
	Score      int    `json:"Score"`
}

type GameSessionLoadResponse struct {
	BoardState string `json:"board_state"`
	Score      int    `json:"Score"`
}

func (g *GameSessionSaveRequest) ToDomain() *model.Session {
	return &model.Session{
		Score:      g.Score,
		BoardState: g.BoardState,
	}
}

func ToGameSessionResponse(m *model.Session) *GameSessionLoadResponse {
	return &GameSessionLoadResponse{
		Score:      m.Score,
		BoardState: m.BoardState,
	}
}
