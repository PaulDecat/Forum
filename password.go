package main

import (
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) PasswordCheckResult {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return PasswordCheckResult{
			Success: false,
			Message: "Invalid password",
		}
	}
	return PasswordCheckResult{
		Success: true,
		Message: "Password is correct",
	}
}
