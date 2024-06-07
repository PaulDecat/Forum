package main
import (
    "errors"
    "fmt"
)

var users = make(map[int]User)
var nextID = 1


func (u *User) Register() error {
    for _, user := range users {
        if user.Email == u.Email || user.Username == u.Username {
            return errors.New("email or username already in use")
        }
    }
    
    u.ID = nextID
    users[u.ID] = *u
    nextID++
    fmt.Println("User registered successfully:", u)
    return nil
}


func (u *User) Login() error {
    for _, user := range users {
        if user.Email == u.Email && user.Password == u.Password {
            fmt.Println("User logged in successfully:", user)
            return nil
        }
    }
    return errors.New("invalid email or password")
}


func (u *User) UpdateProfile() error {
    if existingUser, exists := users[u.ID]; exists {
        existingUser.Email = u.Email
        existingUser.Username = u.Username
        existingUser.Password = u.Password
        users[u.ID] = existingUser
        fmt.Println("User profile updated successfully:", existingUser)
        return nil
    }
    return errors.New("user not found")
}

func (u *User) DeleteAccount() error {
        if _, exists := users[u.ID]; exists {
            delete(users, u.ID)
            fmt.Println("User account deleted successfully:", u)
            return nil
        }
        return errors.New("user not found")
    }
