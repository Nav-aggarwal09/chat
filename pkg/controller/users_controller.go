package controller

import (
	"bytes"
	"fmt"
	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// CreateUser creates a new user
func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nentered CreateUser handler")
	buff := new(bytes.Buffer)
	buff.ReadFrom(r.Body)
	createUserStr := buff.String()
	log.Debug("converted request payload to string: ", createUserStr)
	id, err := models.CreateUser(createUserStr)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	// TODO: Create a New User
	helpers.RespondJSON(w, id)
}
