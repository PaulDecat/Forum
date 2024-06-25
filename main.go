package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialisation de la base de données
	var err error
	db, err = initDB("./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Gestionnaire de fichiers statiques
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Gestionnaires de routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/signup", signUpHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/createP", createPostHandler)
	http.HandleFunc("/terms", termsHandler)
	http.HandleFunc("/rgpd", rgpdHandler)
	http.HandleFunc("/addlike", handleAddLike)
	http.HandleFunc("/adddislike", handleAddDislike)
	http.HandleFunc("/submitcom", submitComment)
	http.HandleFunc("/mypage", mypageHandler)
	http.HandleFunc("/parameters", parametersHandler)
	// http.HandleFunc("/deletepost", deletePostHandler)
	// http.HandleFunc("editpost", editPostHandler)

	// Démarrage du serveur
	log.Println("\x1b[33mStarting server at http://localhost:8080\x1b[0m")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Erreur:", err)
	}
}
