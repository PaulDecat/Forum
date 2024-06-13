package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

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

		// Récupère les commentaires pour chaque post
		comments, err := getCommentsByPostID(post.ID)
		if err != nil {
			return nil, err
		}

		post.Comments = comments

		posts = append(posts, post)
	}

	return posts, nil
}

func getCommentsByPostID(postID int) ([]Comment, error) {
	var comments []Comment

	rows, err := db.Query("SELECT ID, Comments, PostID FROM Comment WHERE PostID = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.Comments, &comment.PostID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
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
		category := r.FormValue("category")

		_, err := db.Exec("INSERT INTO Post (Title, Content, UserID, Category, Likes, Dislikes) VALUES (?, ?, ?, ?, 0, 0)", title, content, userID, category, 0, 0)
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

		category := r.FormValue("category")
		content := r.FormValue("content")
		userID := 1 // Pour test

		_, err := db.Exec("INSERT INTO Post (Title, Content, UserID, Category, Likes, Dislikes) VALUES (?, ?, ?, ?, 0, 0)", title, content, userID, category, 0, 0)
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

func handleAddLike(w http.ResponseWriter, r *http.Request) {
	postIDStr := r.FormValue("Like")
	postID, err := strconv.Atoi(postIDStr)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	updateLikes(w, r, db, postID)
	w.WriteHeader(http.StatusOK)
}

func handleAddDislike(w http.ResponseWriter, r *http.Request) {
	postIDStr := r.FormValue("Dislike")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	updateDislikes(w, r, db, postID)
}

func submitComment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		comment := r.FormValue("comments")
		postIDStr := r.FormValue("postID")
		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		_, err = db.Exec("INSERT INTO Comment (Comments, PostID) VALUES (?, ?)", comment, postID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}
