package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// To make it work, UP your DB:
// docker run -d -e POSTGRES_PASSWORD=postgres -p 5435:5432 postgres

func main() {
	db, err := sql.Open(
		"postgres",
		"host=localhost port=5435 user=postgres password=postgres dbname=postgres sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100)
    )
`)
	if err != nil {
		log.Fatal(err)
	}

	// Adding test data
	_, err = db.Exec(`
    INSERT INTO users (name) VALUES
        ('Alice'),
        ('Bob'),
        ('Charlie')
    ON CONFLICT DO NOTHING
`)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// ID: 1, Name: Alice
// ID: 2, Name: Bob
// ID: 3, Name: Charlie
