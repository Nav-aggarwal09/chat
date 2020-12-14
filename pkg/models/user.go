package models

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/challenge/pkg/database"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"os"
)

type User struct {
	ID       int
	Username string `json:"Username"`
	Password string `json:"Password"`
}

//
func CreateUser(createPayload string) (int, error) {
	var newUser User
	var err error
	ID := 0
	err = json.Unmarshal([]byte(createPayload), &newUser)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not unmarshall into struct: %v\n", err)
		return 0, err
	}
	log.Debugf("Unmarshaled user create payload to %+v", newUser)
	_, usrExists := UserExists(newUser.Username)
	if usrExists {
		return 0, errors.New(fmt.Sprintf("user %s already exists. Please choose another", newUser.Username))
	}
	// TODO: encrypt Password w SHA1 & Salt
	err = database.DBCon.QueryRow(`INSERT INTO users (NULL, ?, ?)`, newUser.Username, newUser.Password).Scan(&ID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not insert new user: %v\n", err)
		return 0, err
	}
	return ID, nil
}

// UserExists returns true if given Username exists and false if not
func UserExists(Username string) (*sql.Row, bool) {
	log.Debugf("Checking if user %s exists", Username)
	sqlStmt := `SELECT 1 FROM users WHERE username = ?`
	row := database.DBCon.QueryRow(sqlStmt, Username)
	fmt.Println("Row Query returned: ", row)
	if row != nil {
		return row, true
	} else {
		return nil, false
	}

}

func AuthenticateUser(loginPayload string) (int, string, error) {
	/*
		1. unmarshal
		2. check if user exists
		3. if yes, check password match
		3b. if no, return
		4. generate token
		5. return id & token
	*/
	log.Debug("Entered AuthenticateUser()")
	var newUser User
	var id int
	var actualUsername string
	var actualPassword string
	err := json.Unmarshal([]byte(loginPayload), &newUser)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not unmarshall into struct: %v\n", err)
		return 0, "", err
	}
	log.Debugf("Unmarshaled user login payload to %+v", newUser)
	usrRow, usrExists := UserExists(newUser.Username)
	if !usrExists {
		return 0, "", errors.New(fmt.Sprintf("user %s does not exist. Please ", newUser.Username))
	}
	err = usrRow.Scan(&id, &actualUsername, &actualPassword)
	if err == sql.ErrNoRows {
		return 0, "", errors.New("No rows returned")
	}
	if newUser.Password == actualPassword {
		log.Debug("Passwords matched")
		return id, "token", nil
	}
	return 0, "", errors.New("Invalid password")

}
