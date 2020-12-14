package helpers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// RespondJSON translates an interface to json for response
func RespondJSON(w http.ResponseWriter, resp interface{}) {
	log.Debug("Attempting to send resp in helpers/RespondJSON()")
	retJSON, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(retJSON)
	if err != nil {
		log.Debug("NOT successful write")
	}
	log.Debug("Succesful write")
}
