package data

import (
	"context"
	"database/sql"
	"time"
)

type Vote struct {
	VoterID     int64     `json:"voterId"`
	CandidateID int64     `json:"candidateId"`
	SessionID   int64     `json:"-"`
	CreatedAt   time.Time `json:"createdAt"`
}

type VoteModel struct {
	DB *sql.DB
}

func (m VoteModel) Insert(v *Vote) error {

	query := `INSERT INTO votes (voter_id, candidate_id, session_id)
				 VALUES ($1, $2, $3) RETURNING created_at`

	args := []any{v.VoterID, v.CandidateID, v.SessionID}

	ctx, cancel := context.WithTimeout(context.Background(), THREE_SECONDS)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&v.CreatedAt)
	return err
}
