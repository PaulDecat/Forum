package main

import (
	"log"
	"net/http"
)

func main() {
	var err error
	db, err = initDB("./blog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", redirectToSignUp)
	http.HandleFunc("/signup", signUpHandler)
	http.HandleFunc("/createpost", createPostHandler)
	http.HandleFunc("/submitpost", submitPostHandler)
	http.HandleFunc("/submituser", submitUserHandler)
	http.HandleFunc("/terms", termsHandler)
	http.HandleFunc("/rgpd", rgpdHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
