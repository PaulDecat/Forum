package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./front/index.html")
}

func createPHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./front/createp.html")
}

func lookfsubHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./front/lookfsub.html")
}

func main() {
	fileServer := http.FileServer(http.Dir("./front"))
	http.Handle("/", fileServer)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/createP", createPHandler)
	http.HandleFunc("/lookfsub", lookfsubHandler)

	fmt.Printf("Starting server at http://localhost:8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erreur:", err)
	}
}
