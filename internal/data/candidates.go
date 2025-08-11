package data

import (
	"context"
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

func (m CandidateModel) Insert(c *Candidate) error {
	query := `INSERT INTO candidates (name, image_url, user_id, session_id, created_at)
				 VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

	args := []any{c.Name, c.ImageURL, c.UserID, c.SessionID, c.CreatedAt}

	ctx, cancel := context.WithTimeout(context.Background(), THREE_SECONDS)

	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&c.ID, &c.CreatedAt)

}
