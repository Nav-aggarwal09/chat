package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type Login struct {
	Token string
	Id    int
}
