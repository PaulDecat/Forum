package main

import (
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Récupère tous les posts depuis la base de données
	posts, err := getAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Charge le modèle HTML
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Passe les posts au modèle HTML pour le rendu
	err = tmpl.Execute(w, posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Récupérer tous les posts depuis la db
func getAllPosts() ([]Post, error) {
	var posts []Post

	// Exécute une requête SQL pour récupérer tous les posts
	rows, err := db.Query("SELECT ID, Title, Content, UserID, Category, Likes FROM Post")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parcourt les lignes résultantes et ajoute les posts à la liste
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID, &post.Category, &post.Likes)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/signup.html")
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "./templates/createP.html")
		return
	} else if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")
		userID := 1 // Pour test

		_, err := db.Exec("INSERT INTO Post (Title, Content, UserID, Category, Likes) VALUES (?, ?, ?, ?, 0)", title, content, userID, "General")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
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
		userID := 1 

		_, err := db.Exec("INSERT INTO Post (Title, Content, UserID, Category, Likes) VALUES (?, ?, ?, ?, 0)", title, content, userID, "General")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}

func termsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/terms.html")
}

func rgpdHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/rgpd.html")
}

func mypageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/mypage.html")
}

func parametersHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/parameters.html")
}

func updateCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		content := r.FormValue("content")

		_, err := db.Exec("UPDATE Comment SET Content = ? WHERE ID = ?", content, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}

func deleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")

		_, err := db.Exec("DELETE FROM Comment WHERE ID = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}

