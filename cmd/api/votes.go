package main

import "net/http"

func (app *application) registerVote(w http.ResponseWriter, r *http.Request) {

	app.logger.Info("Register vote")

}
