package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = createTables(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	createUserTable := `
    CREATE TABLE IF NOT EXISTS User (
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
        Email TEXT,
        Username TEXT,
        Password TEXT
    );`

	createPostTable := `
    CREATE TABLE IF NOT EXISTS Post (
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
        Category TEXT,
        Title TEXT,
        Content TEXT,
        UserID INTEGER,
        Likes INTEGER,
        Dislikes INTEGER,
        FOREIGN KEY(UserID) REFERENCES User(ID)
    );`

	createCommentTable := `
    CREATE TABLE IF NOT EXISTS Comment (
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
        PostID INTEGER,
        Comments TEXT,
        FOREIGN KEY(PostID) REFERENCES Post(ID)
    );`

	_, err := db.Exec(createUserTable)
	if err != nil {
		log.Printf("\x1b[31mError creating User table: %v\x1b[0m", err)
		return err
	}
	log.Println("\x1b[33mUser table created successfully or already exists\x1b[0m")

	_, err = db.Exec(createPostTable)
	if err != nil {
		log.Printf("\x1b[31mError creating Post table: %v\x1b[0m", err)
		return err
	}
	log.Println("\x1b[33mPost table created successfully or already exists\x1b[0m")

	_, err = db.Exec(createCommentTable)
	if err != nil {
		log.Printf("\x1b[31Error creating Post table: %v\x1b[0m", err)
		return err
	}
	log.Println("\x1b[33mComment table created successfully or already exists\x1b[0m")

	return nil
}
