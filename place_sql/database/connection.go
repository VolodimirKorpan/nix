package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // driver mysql
)

// DB ...
var DB *sql.DB

// Connect ...
func Connect() {
	connection, err := sql.Open("mysql", "root:golang@/nix_education")
	if err != nil {
		panic("Could not connect to the database")
	}

	DB = connection

}