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
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), requestHandler()))

}

func home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/crossphoton/IIITR-SERVER", http.StatusTemporaryRedirect)
}

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
