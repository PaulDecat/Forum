package main

type User struct {
    ID       int    `json:"id"`
    Email    string `json:"email"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type Post struct {
    ID       int    `json:"id"`
    Title    string `json:"title"`
    Content  string `json:"content"`
    UserID   int    `json:"user_id"`
    Category string `json:"category"`
}

type Comment struct {
    ID      int    `json:"id"`
    Content string `json:"content"`
    PostID  int    `json:"post_id"`
    UserID  int    `json:"user_id"`
}