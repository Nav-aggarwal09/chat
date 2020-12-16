package controller

import (
	"net/http"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
)

// Check returns the health of the service and DB
func (h Handler) Check(w http.ResponseWriter, r *http.Request) {
	// only respond with ok for now.
	helpers.RespondJSON(w, models.Health{
		"ok",
	})

}
