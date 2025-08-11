package main

import (
	"fmt"
	"net/http"

	"github.com/kaahvote/backend-service-api/internal/data"
)

func (app *application) postSessionCandidatesHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name   string `json:"name"`
		UserID int64  `json:"userId"`
	}

	if input.UserID == 0 {
		// TODO: Create an anonymous user here
		input.UserID = 2
	}

	session, err := app.getSession(r)
	if err != nil {
		app.handleErrToNotFound(w, r, err)
		return
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	c := &data.Candidate{
		Name:      input.Name,
		UserID:    input.UserID,
		SessionID: session.ID,
	}

	// TODO: add validation for candidate fields

	err = app.models.Candidate.Insert(c)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/sessions/%s/candidates/%d", session.PublicID, c.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"candidate": c}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) getSessionCandidatesHandler(w http.ResponseWriter, r *http.Request) {

}

func (app *application) deleteSessionCandidateHandler(w http.ResponseWriter, r *http.Request) {

}
