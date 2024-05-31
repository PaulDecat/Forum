package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error
	db, err = initDB("./blog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fileServer := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fileServer)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/createP", createPostHandler)
	http.HandleFunc("/mypage", mypageHandler)
	http.HandleFunc("/parameters", parametersHandler)
	http.HandleFunc("/signup", signUpHandler)
	http.HandleFunc("/terms", termsHandler)
	http.HandleFunc("/rgpd", rgpdHandler)
	http.HandleFunc("/submituser", submitUserHandler)

	log.Println("Starting server at http://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
