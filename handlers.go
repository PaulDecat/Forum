package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

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

// Récupérer tous les utilisateurs depuis la db
func getAllUsers() ([]User, error) {
	var users []User

	rows, err := db.Query("SELECT ID, Email, Username, Password FROM User")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Hacher le mot de passe avant de l'insérer dans la base de données
		hashedPassword, err := hashPassword(password)
		if err != nil {
			http.Error(w, "Could not hash password", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO User (Email, Username, Password) VALUES (?, ?, ?)", email, username, hashedPassword)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}

	http.ServeFile(w, r, "./templates/signup.html")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		var hashedPassword string
		var userID int

		// Ajout de logs pour vérifier les données entrantes
		log.Printf("Email: %s, Password: %s\n", email, password)

		err := db.QueryRow("SELECT ID, Password FROM User WHERE Email = ?", email).Scan(&userID, &hashedPassword)
		if err != nil {
			log.Printf("Error querying database: %v\n", err)
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// Ajout de logs pour vérifier les données extraites de la base de données
		log.Printf("UserID: %d, Hashed Password: %s\n", userID, hashedPassword)

		checkResult := checkPasswordHash(password, hashedPassword)
		if !checkResult.Success {
			http.Error(w, checkResult.Message, http.StatusUnauthorized)
			return
		}

		session, _ := store.Get(r, "session")
		session.Values["user_id"] = userID
		session.Save(r, w)

		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}

	http.ServeFile(w, r, "./templates/login.html")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	userID := session.Values["userID"]

	if userID == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	var username string
	err := db.QueryRow("SELECT username FROM users WHERE id = ?", userID).Scan(&username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Username string
	}{
		Username: username,
	}

	t, err := template.ParseFiles("./templates/logout.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["userID"] = nil
	session.Save(r, w)

	t.Execute(w, data)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	userID, ok := session.Values["user_id"].(int)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "./templates/createP.html")
		return
	} else if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")

		_, err := db.Exec("INSERT INTO Post (Title, Content, UserID, Category, Likes) VALUES (?, ?, ?, ?, 0)", title, content, userID, "General")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}

func mypageHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	userID := session.Values["userID"]

	if userID == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	_, err := db.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["userID"] = nil
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func termsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/terms.html")
}

func rgpdHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/rgpd.html")
}

func parametersHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/parameters.html")
}
