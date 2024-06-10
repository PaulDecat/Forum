package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func getPostID() {
	var id int
	err := db.QueryRow("SELECT id FROM Post ORDER BY id DESC LIMIT 1").Scan(&id)
	if err != nil {
		fmt.Println("Error getting postId from database")
	}

}

func addlike(postId int) error {
	al, err := db.Prepare("UPDATE Post SET likes = likes + 1 WHERE id =?")
	if err != nil {
		return err
	}
	_, err = al.Exec(postId)
	if err != nil {
		return err
	}
	return nil
}
