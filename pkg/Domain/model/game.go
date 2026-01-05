package model

import "time"

type Session struct {
	UserID       string
	Score        int
	BoardState   string
	LastPlayedAt time.Time
}
