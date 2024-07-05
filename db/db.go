package db

import (
	"database/sql"
	"fmt"
	"os"
)

func Connect() {
	dbName := "file:../coffee.db"
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open database %s", err)
		os.Exit(1)
	}
	defer db.Close()
}
