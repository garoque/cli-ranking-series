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

	db, err = sql.Open("sqlite3", "./cli-conteudos.db")
	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateTable() {
	query := `CREATE TABLE IF NOT EXISTS contents (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		definition TEXT,
		category TEXT
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("created table contents")
}

func InsertNote(word string, definition string, category string) {
	query := `INSERT INTO contents(title, definition, category) VALUES (?, ?, ?)`
	_, err := db.Exec(query, word, definition, category)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted contents note successfully")
}

func DisplayAllNotes() {
	row, err := db.Query("SELECT * FROM contents ORDER BY title")
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var id int
		var title string
		var definition string
		var category string
		row.Scan(&id, &title, &definition, &category)
		log.Println("[", category, "] ", title, "—", definition)
	}
}
