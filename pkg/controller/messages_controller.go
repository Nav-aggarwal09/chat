package controller

import (
	"bytes"
	"github.com/challenge/pkg/helpers"
	"github.com/challenge/pkg/models"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type PostMsg struct {
	Token string

}

// SendMessage send a message from one user to another
func (h Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	// converting io.ReadCloser to string to pass pass around to server
	buff := new(bytes.Buffer)
	buff.ReadFrom(r.Body)
	sendMessageStr := buff.String()
	log.Debug("converted request payload to string: ", sendMessageStr)


	id, timeStamp := models.MessengerMain(r.Header.Get("token"), sendMessageStr)
	helpers.RespondJSON(w, models.ResponseMessage{ID: id, Timestamp: timeStamp})
}

// GetMessages get the messages from the logged user to a recipient
func (h Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	// TODO: Retrieve list of Messages
	helpers.RespondJSON(w, []*models.Message{{}})
}
