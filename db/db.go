package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}
type Coffee struct {
	Name    string
	Roaster string
	Origins string
}

type connectError struct {
	message string
}

func (c *connectError) Error() string {
	return c.message
}

func Connect() (Database, error) {
	dbName := "file:./coffee.db"
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return Database{}, &connectError{message: "Cannot Connect to Database"}
	}
	return Database{db: db}, nil
}

func (c *Database) GetAllCoffees() ([]Coffee, error) {
	rows, err := c.db.Query("SELECT name, roaster, origins from coffee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := make([]Coffee, 0)

	for rows.Next() {
		coffee := Coffee{}
		err = rows.Scan(&coffee.Name, &coffee.Roaster, &coffee.Origins)
		if err != nil {
			return nil, err
		}
		data = append(data, coffee)
	}
	return data, nil
}

func (c *Database) CreateNewCoffee(name string, roaster string, origins string) (int, error) {
	res, err := c.db.Exec("INSERT INTO coffee VALUES(null, ?, ?, ?)", name, roaster, origins)
	if err != nil {
		return 0, err
	}
	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, nil
	}
	return int(id), nil
}
