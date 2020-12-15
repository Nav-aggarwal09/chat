package controller

import (
	"bytes"
	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type createdUser struct {
	Id int
}

// CreateUser creates a new user
func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	buff := new(bytes.Buffer)
	buff.ReadFrom(r.Body)
	createUserStr := buff.String()
	log.Debug("converted request payload to string: ", createUserStr)
	id, err := models.CreateUser(createUserStr)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	helpers.RespondJSON(w, createdUser{Id: id})
}
