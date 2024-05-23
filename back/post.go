package main

type Post struct {
    ID       int    `json:"id"`
    Title    string `json:"title"`
    Content  string `json:"content"`
    UserID   int    `json:"user_id"`
    Category string `json:"category"`
    Like     int    `json:"like"`
}

func (p *Post) CreatePost() {
    
}

func (p *Post) UpdatePost() {
 
}

func (p *Post) DeletePost() {
   
}

func (p *Post) LikePost() {
 
}