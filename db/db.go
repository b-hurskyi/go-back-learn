package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("pgx", "postgres://myuser:mypassword@localhost:5434/mydb?sslmode=disable")

	if err != nil {
		fmt.Println(err)
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			date_time TIMESTAMP NOT NULL,
			user_id INTEGER
		);
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table.")
	}
}
