package data

import (
	"database/sql"
	"time"
)

type Candidate struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	ImageURL  string    `json:"imageUrl"`
	UserID    int64     `json:"userId"`
	SessionID int64     `json:"sessionId"`
	CreatedAt time.Time `json:"createdAt"`
}

type CandidateModel struct {
	DB *sql.DB
}
