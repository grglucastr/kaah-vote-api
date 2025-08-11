package main

import (
	"net/http"

	"github.com/kaahvote/backend-service-api/internal/data"
	"github.com/kaahvote/backend-service-api/internal/validator"
)

func (app *application) getVotingPoliciesHandler(w http.ResponseWriter, r *http.Request) {

	qs := r.URL.Query()

	name := app.readString(qs, "name", "")
	createdFrom := app.readDate(qs, "createdFrom")
	createdTo := app.readDate(qs, "createdTo")

	v := validator.New()
	page := app.readInt(qs, "page", 1, v)
	pageSize := app.readInt(qs, "pageSize", 5, v)
	sort := app.readString(qs, "sort", "createdAt")

	filters := data.SettingsFilters{
		Name:          name,
		CreatedAtFrom: createdFrom,
		CreatedAtTo:   createdTo,
		Filters: data.Filters{
			Page:         page,
			PageSize:     pageSize,
			Sort:         sort,
			SortSafeList: []string{"name", "createdAt", "-name", "createdAt"},
		},
	}

	policies, metadata, err := app.models.VotingPolicy.ListFiltering(filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"metadata": metadata, "policies": policies}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
