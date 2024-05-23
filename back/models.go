package main

type User struct {
    ID       int    `json:"id"`
    Email    string `json:"email"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func (u *User) Register() {
    
}

func (u *User) Login() {
   
}

func (u *User) UpdateProfile() {
  
}

func (u *User) DeleteAccount() {
   
}


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

type Comment struct {
    ID      int    `json:"id"`
    Content string `json:"content"`
    PostID  int    `json:"post_id"`
    UserID  int    `json:"user_id"`
}

func (c *Comment) CreateComment() {
   
}

func (c *Comment) UpdateComment() {
  
}

func (c *Comment) DeleteComment() {

}