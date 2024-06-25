package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/http"
	"time"
)

// Générer un ID de session sécurisé
func generateSessionID() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// Définir un cookie de session
func setSessionCookie(w http.ResponseWriter, sessionID string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(1 * time.Hour), // Date d'expiration du cookie (1 heure)
	})
}

// Récupérer l'ID de session à partir d'un cookie
func getSessionID(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// Supprimer un cookie de session
func clearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id", // Nom du cookie
		Value:    "",           // Valeur vide pour effacer le cookie
		Path:     "/",          // Chemin pour lequel le cookie est valide
		HttpOnly: true,         // Empêche l'accès au cookie via JavaScript (sécurité)
		MaxAge:   -1,           // MaxAge négatif pour indiquer que le cookie est expiré
	})
}

// variable sessions en mémoire
var sessionStore = map[string]int{}

// Enregistrer l'ID de session & l'ID utilisateur ( login )
func saveSession(sessionID string, userID int) {
	sessionStore[sessionID] = userID
}

// Obtenir l'ID utilisateur à partir de l'ID de session
func getUserIDFromSession(sessionID string) (int, error) {
	userID, exists := sessionStore[sessionID]
	if !exists {
		return 0, errors.New("session not found")
	}
	return userID, nil
}

// Supprimer une session de la mémoire( déconnexion )
func removeSession(sessionID string) {
	delete(sessionStore, sessionID)
}
