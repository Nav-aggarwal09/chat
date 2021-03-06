package controller

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
)

// Login authenticates a user and returns a token
func (h Handler) Login(w http.ResponseWriter, r *http.Request) {

	// converting io.ReadCloser to string to pass pass around to server
	buff := new(bytes.Buffer)
	buff.ReadFrom(r.Body)
	LoginStr := buff.String()
	log.Debug("converted request payload to string: ", LoginStr)

	id, token, err := models.AuthenticateUser(LoginStr)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	helpers.RespondJSON(w, models.Login{Id: id, Token: token})
}
