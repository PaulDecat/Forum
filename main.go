package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error
	db, err = initDB("./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/signup", signUpHandler)
	http.HandleFunc("/createP", createPostHandler)
	http.HandleFunc("/submituser", submitUserHandler)
	http.HandleFunc("/submitpost", submitPostHandler)
	http.HandleFunc("/terms", termsHandler)
	http.HandleFunc("/rgpd", rgpdHandler)
	http.HandleFunc("/addlike", handleAddLike)
	http.HandleFunc("/adddislike", handleAddDislike)
	http.HandleFunc("/deletepost", deletePostHandler)
	http.HandleFunc("editpost", editPostHandler)
	http.HandleFunc("/submitcom", submitComment)
	http.HandleFunc("/mypage", mypageHandler)
	http.HandleFunc("/parameters", parametersHandler)

	log.Println("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
