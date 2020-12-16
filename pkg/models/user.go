package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/challenge/pkg/database"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type User struct {
	ID       int
	Username string `json:"Username"`
	Password string `json:"Password"`
}

//CreateUser enters the requesting User credentials into the DB if valid
func CreateUser(createPayload string) (int, error) {
	/*
		1. unmarshal client information from given string
		2. check if username already exists and non-nil credentials
		3. Encrypt password
		4. insert into DB
		5. return ID
	*/
	var newUser User
	var err error
	err = json.Unmarshal([]byte(createPayload), &newUser)
	if err != nil {
		log.Errorf("could not unmarshall into struct: %v\n", err)
		return 0, err
	}
	log.Debugf("Unmarshaled user create payload to %+v", newUser)
	if newUser.Username == "" || newUser.Password == "" {
		return 0, errors.New(fmt.Sprintf("Please give a non-empty username or password"))
	}
	_, usrExists := UserExists(newUser.Username)
	if usrExists {
		return 0, errors.New(fmt.Sprintf("user %s already exists. Please choose another", newUser.Username))
	}

	// TODO: this is where we would encrypt Password w SHA1 & Salt
	addUsrStatement, err := database.DBCon.Prepare(`INSERT INTO Users (username, password) values ($1, $2)`)
	if err != nil {
		log.Errorf("Could not prepare query to insert new user: %v", err)
		return 0, err
	}

	result, err := addUsrStatement.Exec(newUser.Username, newUser.Password)
	if err != nil {
		log.Errorf("Could not execute query to insert new user: %v", err)
		return 0, err
	}
	id, _ := result.LastInsertId()
	log.Info("Successfully created new user with id ", id)
	return int(id), nil

}

// UserExists returns true if given Username exists and false if not
func UserExists(Username string) (User, bool) {
	log.Debugf("Checking if user %s exists", Username)
	var newUser User
	row := database.DBCon.QueryRow(`SELECT id, username, password FROM Users WHERE username = ?`, Username)
	_ = row.Scan(&newUser.ID, &newUser.Username, &newUser.Password)
	log.Debugf("UserExists query returned - ID %d username %s password %s", newUser.ID, newUser.Username, newUser.Password)

	if newUser.Username == "" {
		log.Debug("Username not found")
		return User{}, false
	} else {
		return newUser, true
	}

}

// AuthenticateUser authenticates the requesting User by checking credentials with database
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
	err := json.Unmarshal([]byte(loginPayload), &newUser)
	if err != nil {
		log.Errorf("could not unmarshall into struct: %v", err)
		return 0, "", err
	}
	log.Debugf("Unmarshaled user login payload to %+v", newUser)
	dbUsr, usrExists := UserExists(newUser.Username)
	if !usrExists {
		return 0, "", errors.New(fmt.Sprintf("user %s does not exist", newUser.Username))
	}
	if newUser.Password != dbUsr.Password {
		return 0, "", errors.New("Invalid password")
	}
	log.Debug("Passwords matched")

	token := AddTokenAndTimeStamp(newUser.Username)
	return dbUsr.ID, token, nil
}
