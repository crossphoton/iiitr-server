/**
This application is an extension to iiitr-services/Server

Auth package implements authentication layer for user login.

Currently supported methods
	- Google OAuth 			(for students)


This extension assumes these environment variables to be available (besides from main Server)
	- GOOGLE_CLIENT_ID			- For Google OAuth
	- GOOGLE_CLIENT_SECRET		- For Google OAuth
	- DOMAIN					- For cookie formation
	- JWT_SIGNING_KEY			- For JWT formation
*/
package auth

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Handler for handling auth endpoints and extension activation
func Handler(r *mux.Router, database *gorm.DB) {

	db = database
	db.AutoMigrate(Student{})

	// Google OAuth
	r.HandleFunc("/google", googleOAuthRedirect).Methods("GET")
	r.HandleFunc("/google/callback", googleOAuthCallback).Methods("GET")
}
