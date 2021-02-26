package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"example.com/auth"
	"example.com/studentdata"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func requestHandler() *mux.Router {
	handler := mux.NewRouter()

	handler.HandleFunc("/", home).Methods("GET")
	auth.Handler(handler.PathPrefix("/auth").Subrouter(), db)
	studentdata.Handler(handler.PathPrefix("/studentdata").Subrouter(), db)
	return handler
}

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		db.Close()
		fmt.Println("\nShutting down server.\nBye Bye...")
		os.Exit(0)
	}()

	dbInit()

	fmt.Println("Starting Up server on port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), requestHandler()))

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<a href=\"/auth/google\">Login</a>")
}

func dbInit() {
	addr := os.Getenv("DB_URL")
	db, err = gorm.Open("postgres", addr)

	if err != nil {
		log.Fatal(err)
		log.Fatal("DB Error")
	}

	db.AutoMigrate(&studentdata.Student{})
	db.AutoMigrate(&studentdata.AIMSAcademicData{})
}
