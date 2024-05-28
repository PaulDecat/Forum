package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

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
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./blog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTables()

	http.HandleFunc("/", redirectToSignUp)
	http.HandleFunc("/signup", signUpHandler)
	http.HandleFunc("/createpost", createPostHandler)
	http.HandleFunc("/submituser", submitUserHandler)
	http.HandleFunc("/submitpost", submitPostHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createTables() {
	createUserTable := `
    CREATE TABLE IF NOT EXISTS User (
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
        Email TEXT,
        Username TEXT,
        Password TEXT
    );`

	createPostTable := `
    CREATE TABLE IF NOT EXISTS Post (
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
        Title TEXT,
        Content TEXT,
        UserID INTEGER,
        Category TEXT,
        Likes INTEGER,
        FOREIGN KEY(UserID) REFERENCES User(ID)
    );`

	_, err := db.Exec(createUserTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(createPostTable)
	if err != nil {
		log.Fatal(err)
	}
}

func redirectToSignUp(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/signup", http.StatusSeeOther)
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/signup.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/createpost.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func submitUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")

		_, err := db.Exec("INSERT INTO User (Email, Username, Password) VALUES (?, ?, ?)", email, username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	}
}

func submitPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")
		userID := 1 // Pour test

		_, err := db.Exec("INSERT INTO Post (Title, Content, UserID, Category, Likes) VALUES (?, ?, ?, ?, 0)", title, content, userID, "General")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/createpost", http.StatusSeeOther)
	}
}
