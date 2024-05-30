package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./blog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTables()

	fileServer := http.FileServer(http.Dir("./front"))
	http.Handle("/", fileServer)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/createP", createPostHandler)
	http.HandleFunc("/mypage", mypageHandler)
	http.HandleFunc("/parameters", parametersHandler)
	http.HandleFunc("/signup", signUpHandler)
	http.HandleFunc("/terms", termsHandler)
	http.HandleFunc("/rgpd", rgpdHandler)
	http.HandleFunc("/submituser", submitUserHandler)

	fmt.Printf("Starting server at http://localhost:8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erreur:", err)
	}
}
