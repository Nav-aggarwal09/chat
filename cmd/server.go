package main

import (
	"database/sql"
	"fmt"
	"github.com/challenge/pkg/auth"
	"github.com/challenge/pkg/controller"
	"github.com/challenge/pkg/database"
	"os"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	ServerPort       = "8080"
	CheckEndpoint    = "/check"
	UsersEndpoint    = "/users"
	LoginEndpoint    = "/login"
	MessagesEndpoint = "/messages"
)


func main() {
	var err error
	// using a logging framework to keep track of server errors and actions
	log.SetLevel(log.DebugLevel)
	log.Info("starting program...")

	db, err := sql.Open("sqlite3", "./chat.db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open database chat.db: %v\n", err)
		os.Exit(1)
	}

	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS "Users" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"username" text NOT NULL UNIQUE ,
		"password" text NOT NULL,
		"token" text,
		"timestamp" text);`)
	_, err = stmt.Exec()
	if err != nil {
		log.Error("Error creating users table")
		fmt.Println(err)
		os.Exit(1)
	}
	database.DBCon = db
	log.Info("opened sql connection")
	h := controller.Handler{}

	// Configure endpoints
	// Health
	http.HandleFunc(CheckEndpoint, func(w http.ResponseWriter, r *http.Request) {
		log.Debug("connection hit /check endpoint")
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		h.Check(w, r)
	})

	// Users
	http.HandleFunc(UsersEndpoint, func(w http.ResponseWriter, r *http.Request) {
		log.Debug("connection hit /users endpoint")
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
		h.CreateUser(w, r)
	})

	// Auth
	http.HandleFunc(LoginEndpoint, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}

		h.Login(w, r)
	})

	// Messages
	http.HandleFunc(MessagesEndpoint, auth.ValidateUser(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetMessages(w, r)
		case http.MethodPost:
			h.SendMessage(w, r)
		default:
			http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
			return
		}
	}))

	// Start server
	log.Println("Server started at port " + ServerPort)
	log.Fatal(http.ListenAndServe(":"+ServerPort, nil))
}
