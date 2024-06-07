package main

import (
    "errors"
    "fmt"
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
    if existingComment, exists := comments[c.ID]; exists {
        existingComment.Content = c.Content
        comments[c.ID] = existingComment
        fmt.Println("Commentaire mis à jour avec succès :", existingComment)
        return nil
    }
    return errors.New("commentaire non trouvé")
}

func (c *Comment) DeleteComment() error {
    if _, exists := comments[c.ID]; exists {
        delete(comments, c.ID)
        fmt.Println("Commentaire supprimé avec succès :", c)
        return nil
    }
    return errors.New("commentaire non trouvé")
}
