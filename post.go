package main

import(
    "errors"
    "fmt"
)






var posts = make(map[int]Post)
var nextPostID = 1


func (p *Post) CreatePost() error {
    p.ID = nextPostID
    posts[p.ID] = *p
    nextPostID++
    fmt.Println("Publication créée avec succès :", p)
    return nil
}

func (p *Post) UpdatePost() error {
    if existingPost, exists := posts[p.ID]; exists {
        existingPost.Title = p.Title
        existingPost.Content = p.Content
        existingPost.Category = p.Category
        posts[p.ID] = existingPost
        fmt.Println("Publication mise à jour avec succès :", existingPost)
        return nil
    }
    return errors.New("publication non trouvée")
}

func (p *Post) DeletePost() error {
    if _, exists := posts[p.ID]; exists {
        delete(posts, p.ID)
        fmt.Println("Publication supprimée avec succès :", p)
        return nil
    }
    return errors.New("publication non trouvée")
}

func (p *Post) LikePost() error {
    if existingPost, exists := posts[p.ID]; exists {
        existingPost.Likes++
        posts[p.ID] = existingPost
        fmt.Println("Publication likée avec succès :", existingPost)
        return nil
    }
    return errors.New("publication non trouvée")
}