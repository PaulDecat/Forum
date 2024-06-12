package main

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := getAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Récupérer tous les posts depuis la db
func getAllPosts() ([]Post, error) {
	var posts []Post

	rows, err := db.Query("SELECT ID, Title, Content, UserID, Category, Likes FROM Post")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

		log.Printf("\x1b[32mAccount created: Username: %s, Email: %s\x1b[0m\n", username, email)

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

		log.Printf("New Entry in login form: Email: %s, Password: %s\n", email, password)

		err := db.QueryRow("SELECT ID, Password FROM User WHERE Email = ?", email).Scan(&userID, &hashedPassword)
		if err != nil {
			log.Printf("\x1b[31mError querying database: %v\x1b[0m\n", err)
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		checkResult := checkPasswordHash(password, hashedPassword)
		if !checkResult.Success {
			http.Error(w, checkResult.Message, http.StatusUnauthorized)
			return
		}

		sessionID, err := generateSessionID()
		if err != nil {
			http.Error(w, "Could not create session", http.StatusInternalServerError)
			return
		}

		setSessionCookie(w, sessionID)
		saveSession(sessionID, userID)

		log.Printf("\x1b[34mUser logged in: UserID: %d, Email: %s, Hashed Password: %s\x1b[0m\n", userID, email, hashedPassword)

		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}

	http.ServeFile(w, r, "./templates/login.html")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, err := getSessionID(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	removeSession(sessionID)
	clearSessionCookie(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, err := getSessionID(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	userID, err := getUserIDFromSession(sessionID)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "./templates/createP.html")
		return
	} else if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")

		result, err := db.Exec("INSERT INTO Post (Title, Content, UserID, Category, Likes) VALUES (?, ?, ?, ?, 0)", title, content, userID, "General")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		postID, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Printf("Post created - Title: %s, Post ID: %d, User ID: %d", title, postID, userID)

		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}

func mypageHandler(w http.ResponseWriter, r *http.Request) {
	// Vérifiez si la méthode est POST
	if r.Method == http.MethodPost {
		sessionID, err := getSessionID(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		userID, err := getUserIDFromSession(sessionID)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		_, err = db.Exec("DELETE FROM User WHERE ID = ?", userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		removeSession(sessionID)
		clearSessionCookie(w)

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Si la méthode est GET, affichez la page du profil
	sessionID, err := getSessionID(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	userID, err := getUserIDFromSession(sessionID)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	var username string
	err = db.QueryRow("SELECT Username FROM User WHERE ID = ?", userID).Scan(&username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Username string
	}{
		Username: username,
	}

	tmpl, err := template.ParseFiles("./templates/mypage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("\x1b[31mAccount deleted: UserID: %d, Username: %s\x1b[0m\n", userID, username)

	tmpl.Execute(w, data)
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
