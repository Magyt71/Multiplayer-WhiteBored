package main

import (
	"database/sql"
	"log/slog"

	// مكتبة حديثة للـ Logging في Go
	_ "github.com/mattn/go-sqlite3" // الـ Driver الخاص بـ SQLite
)

var db *sql.DB

func initdb() {
	var err error
	db, err = sql.Open("sqlite3", "./Whitebored.DB")
	if err != nil {
		slog.Error("Failed connecting to the database", "error", err)
		panic(err)
	}

	queries := `
	Create Table If Not EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS rooms(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS strokes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		room_id TEXT,
		data TEXT
	);
	`
	_, err = db.Exec(queries)
	if err != nil {
		slog.Error("Failed to Create the Tables", "error", err)
		panic(err)
	}
	slog.Info("Done Creating the Tables")
}
