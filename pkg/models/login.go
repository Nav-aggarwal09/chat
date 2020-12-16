package models

import (
	"fmt"
	"github.com/challenge/pkg/database"
	_ "github.com/mattn/go-sqlite3"
	"crypto/rand"
	log "github.com/sirupsen/logrus"
	"time"
)

type Login struct {
	Token string
	Id    int
}

// generateToken generates a random 8 character token
func generateToken() string {
	newToken := make([]byte, 4)
	rand.Read(newToken)
	return fmt.Sprintf("%x", newToken)
}

//AddTokenAndTimeStamp generates token and timestamp and inserts them into database
func AddTokenAndTimeStamp(usrNme string) string {
	tok := generateToken()
	timeStmp := time.Now().Format(time.RFC822)
	addStmt, err := database.DBCon.Prepare(`UPDATE Users SET token = ?, timestamp = ? WHERE username = ?`)
	if err != nil {
		log.Errorf("Could not prepare query to insert token/timestamp: %v", err)
		return ""
	}
	_, err = addStmt.Exec(tok, timeStmp, usrNme)
	if err != nil {
		log.Errorf("Could not execute query to insert token/timestamp: %v", err)
		return ""
	}
	return tok
}
