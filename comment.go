package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Comment struct {
	ID      int
	Content string
	PostID  int
	UserID  int
}

var comments = make(map[int]Comment)
var nextCommentID = 1


func (c *Comment) CreateComment() error {
	c.ID = nextCommentID
	comments[c.ID] = *c
	nextCommentID++
	fmt.Println("Commentaire créé avec succès :", c)
	return nil
}

func (c *Comment) UpdateComment() error {
	_, err := db.Exec("UPDATE Comment SET Content = ? WHERE ID = ?", c.Content, c.ID)
	if err != nil {
		return err
	}
       fmt.Println("Commentaire mis à jour avec succès:", c)
	return nil
}
func (c *Comment) DeleteComment() error {
	_, err := db.Exec("DELETE FROM Comment WHERE ID = ?", c.ID)
	if err != nil {
		return err
	}
	fmt.Println("Commentaire supprimé avec succès:", c)
	return nil
}
