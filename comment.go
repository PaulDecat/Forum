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

func (c *Comment) UpdateComment() {
        result, err := db.Exec("UPDATE Comment SET Content = ? WHERE ID = ?", c.Content, c.ID)
        if err != nil {
            fmt.Println("Erreur lors de la mise à jour du commentaire:", err)
            return
        }
    
        rowsAffected, err := result.RowsAffected()
        if err != nil {
            fmt.Println("Erreur lors de la récupération des lignes affectées:", err)
            return
        }
    
        if rowsAffected == 0 {
            fmt.Println("Aucun commentaire trouvé avec cet ID pour la mise à jour")
            return
        }
        fmt.Println("Commentaire mis à jour avec succès :", c)
    }

func (c *Comment) DeleteComment() {
	_, err := db.Exec("DELETE FROM Comment WHERE ID = ?", c.ID)
	if err != nil {
		fmt.Println("Erreur lors de la suppression du commentaire:", err)
		return
    }
}
