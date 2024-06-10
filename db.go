package main

import (
	"database/sql"
	"log"
	"net/http"

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

	_, err := db.Exec(createUserTable)
	if err != nil {
		log.Printf("Error creating User table: %v", err)
		return err
	}
	log.Println("User table created successfully or already exists")

	_, err = db.Exec(createPostTable)
	if err != nil {
		log.Printf("Error creating Post table: %v", err)
		return err
	}
	log.Println("Post table created successfully or already exists")

	return nil
}

func updateLikes(w http.ResponseWriter, r *http.Request, db *sql.DB, postID int) {
	updateMyLikes := `UPDATE Post SET Likes = Likes + 1 WHERE ID =?;`
	_, err := db.Exec(updateMyLikes, postID)

	println("dffsdfsd")

	if err != nil {
		log.Printf("Error updating likes for post ID %d: %v", postID, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func updateDislikes(w http.ResponseWriter, r *http.Request, db *sql.DB, postID int) {
	updateMyDislikes := `UPDATE Post SET Likes = Likes - 1 WHERE ID =?;`
	_, err := db.Exec(updateMyDislikes, postID)

	if err != nil {
		log.Printf("Error updating dislikes for post ID %d: %v", postID, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
