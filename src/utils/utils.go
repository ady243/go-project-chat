package utils

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func InitDatabase() {
	conn, err := sql.Open("sqlite3", "./chat.db")
	if err != nil {
		log.Fatal(err)
	}

	db = conn

	createTables()
}

func CloseDatabase() {
	db.Close()
}

func createTables() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Users (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Username TEXT NOT NULL,
			Password TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Rooms (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Name TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Messages (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			RoomID INTEGER NOT NULL,
			UserID INTEGER NOT NULL,
			Content TEXT NOT NULL,
			CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}
