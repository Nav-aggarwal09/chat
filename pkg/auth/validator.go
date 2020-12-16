package auth

import (
	"bytes"
	"encoding/json"
	"github.com/challenge/pkg/database"
	"github.com/challenge/pkg/models"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// ValidateUser checks for a token and validates it
// before allowing the method to execute
func ValidateUser(_ http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: validate token
		//     1. get user's token from DB and compare with given token
		//     2. if yes, check if given token is old
		//     3. if not old, do nothing (?)
		// 		if anything fails, http.Error()

		givenTok := r.Header.Get("token")
		var dbToken string
		var dbTimestamp time.Time

		buff := new(bytes.Buffer)
		buff.ReadFrom(r.Body)
		sendMessageStr := buff.String()
		var newMessage models.Message
		var err error

		err = json.Unmarshal([]byte(sendMessageStr), &newMessage)
		if err != nil {
			log.Errorf("could not unmarshall into struct: %v\n", err)
			http.Error(w, "Unable to marshal struct", http.StatusInternalServerError)
		}
		row := database.DBCon.QueryRow(`SELECT token, timestamp FROM Users WHERE id = ?`, newMessage.Sender)
		_ = row.Scan(&dbToken, &dbTimestamp)

		if givenTok != dbToken {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		}
		if dbTimestamp.After(dbTimestamp.Add(time.Hour * 3)) {
			// if invalid, null it out in the DB and force user to get new token by logging in again
			addUsrStatement, _ := database.DBCon.Prepare(`UPDATE Users SET token = NULL, timestamp = NULL WHERE token = dbtoken`)
			addUsrStatement.Exec()
			http.Error(w, "Expired token", http.StatusUnauthorized)
		}
	}
}
