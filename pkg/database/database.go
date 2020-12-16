package database

import "database/sql"
// this package exists to avoid cyclic dependencies

var (
	// DBCon is the connection handle
	// for the database
	DBCon *sql.DB
)

// TODO: add functions to handle queries like INSERT and SELECT
