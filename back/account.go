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