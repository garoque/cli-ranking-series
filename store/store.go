package store

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./cli-ranking-series.db")
	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS series (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		position TEXT
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("created table series")
}

func Insert(title, position string) {
	query := `INSERT INTO series (title, position) VALUES (?, ?)`
	_, err := db.Exec(query, title, position)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted successfully")
}

func DisplayAllNotes() {
	row, err := db.Query("SELECT * FROM series ORDER BY position")
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var title string
		var position int
		row.Scan(&title, &position)
		fmt.Println("[", position, "] ", title, "—", position)
	}
}
