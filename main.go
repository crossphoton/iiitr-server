/*
This is the package for the main Server of the services. Services are treated as extensions to the server.

The SERVER expects the following environment variables to be present
	- PORT 		- Port to serve on
	- DB_URL	- URL of a Postgres protocol supporting database

*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/crossphoton/iiitr-server/auth"
	"github.com/crossphoton/iiitr-server/studentdata"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

// requestHandler make subrouters for extensions
func requestHandler() *mux.Router {
	handler := mux.NewRouter()

	handler.HandleFunc("/", home).Methods("GET")
	auth.Handler(handler.PathPrefix("/auth").Subrouter(), db)
	studentdata.Handler(handler.PathPrefix("/studentdata").Subrouter(), db)
	return handler
}

func main() {

	c := make(chan os.Signal, 1) /* Just a fun thing to do */
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		db.Close()
		fmt.Println("\nShutting down server.\nBye Bye...")
		os.Exit(0)
	}()

	dbInit()

	handler := requestHandler()

	server := http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      handler,
		WriteTimeout: time.Second * 5,
	}

	fmt.Println("Starting Up server on port " + os.Getenv("PORT"))
	log.Fatal(server.ListenAndServe())

}

// Home serves the homepage of the server
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprintln(w, "This is home. Consider visiting <a target=\"_blank\" href=\"https://www.github.com/crossphoton/iiitr-server\">here</a> and give a star.")
}

// dbInit initializes the database
func dbInit() {
	addr := os.Getenv("DB_URL")
	db, err = gorm.Open("postgres", addr)

	if err != nil {
		log.Fatal(err)
		log.Fatal("DB Error")
	}

	db.AutoMigrate(&auth.Student{})
	db.AutoMigrate(&studentdata.AIMSAcademicData{})
}
