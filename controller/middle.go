package controller

import (
	"log"
	"net/http"
)

// make sure the user has session, or redirect to login page
func middleAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := getSessionUser(r)
		log.Println("middle: ", username)
		if err != nil {
			log.Println("middle gets session error and will redirect to login")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}