package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"time"
)

// Fonction pour générer un ID de session sécurisé
func generateSessionID() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// Fonction pour définir un cookie de session
func setSessionCookie(w http.ResponseWriter, sessionID string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(1 * time.Hour),
	})
}

// Fonction pour récupérer l'ID de session à partir d'un cookie
func getSessionID(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// Fonction pour supprimer un cookie de session
func clearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})
}

// Fonction pour enregistrer l'ID de session et l'ID utilisateur en mémoire
var sessionStore = map[string]int{}

func saveSession(sessionID string, userID int) {
	sessionStore[sessionID] = userID
}

func getUserIDFromSession(sessionID string) (int, error) {
	userID, exists := sessionStore[sessionID]
	if !exists {
		return 0, errors.New("session not found")
	}
	return userID, nil
}

func removeSession(sessionID string) {
	delete(sessionStore, sessionID)
}
