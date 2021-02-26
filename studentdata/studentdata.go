package studentdata

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Handler handles http requests
func Handler(r *mux.Router, database *gorm.DB) {
	db = database
	r.HandleFunc("/updateData", updateData).Methods("POST")
}

func updateData(w http.ResponseWriter, r *http.Request) {
	var receivedData map[string]interface{}

	json.NewDecoder(r.Body).Decode(&receivedData)

	lfl, err := json.Marshal(receivedData["data"].(map[string]interface{}))

	if err != nil {
		log.Println(err)
		return
	}

	AIMSdata := AIMSAcademicData{
		Email: receivedData["email"].(string),
		Data:  string(lfl),
	}

	saveData(AIMSdata)
}
