package main

import (
	"net/http"

	"github.com/kaahvote/backend-service-api/internal/data"
)

func (app *application) registerVote(w http.ResponseWriter, r *http.Request) {

	var input struct {
		CandidateID *int64 `json:"candidateId"`
		VoterID     *int64 `json:"voterId"`
		SessionID   *int64 `json:"sessionID"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	vote := &data.Vote{
		VoterID:     input.VoterID,
		CandidateID: input.CandidateID,
		SessionID:   input.SessionID,
	}

	err = app.models.Votes.Insert(vote)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"vote": *vote}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
