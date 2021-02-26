package auth

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Handler for handling auth
func Handler(r *mux.Router, database *gorm.DB) {

	db = database

	// Google OAuth
	r.HandleFunc("/google", googleOAuthRedirect).Methods("GET")
	r.HandleFunc("/google/callback", googleOAuthCallback).Methods("GET")
}
