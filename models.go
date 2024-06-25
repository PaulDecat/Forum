package main

type User struct {
	ID       int
	Email    string
	Username string
	Password string
}

type Post struct {
	ID       int
	Title    string
	Content  string
	UserID   int
	Category string
	Likes    int
	Dislikes int
	Comments []Comment
}

type Comment struct {
	ID       int
	Comments string
	PostID   int
}

type PageData struct {
	Posts      []Post
	IsLoggedIn bool
}

type PasswordCheckResult struct {
	Success bool
	Message string
}
