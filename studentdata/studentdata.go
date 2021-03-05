package studentdata

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Handler handles http requests
func Handler(r *mux.Router, database *gorm.DB) {
	db = database
	r.HandleFunc("/updateData", updateData).Methods("POST")
}

func updateData(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized")
		return
	}

	var body map[string]interface{}
	json.NewDecoder(r.Body).Decode(&body)

	tokenStatus := verifyJWT(token.Value, body["email"].(string))
	if tokenStatus != true {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized")
		return
	}

	b, err := json.Marshal(body["data"].([]interface{}))
	receivedData := string(b)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Input", 400)
		return
	}

	AIMSdata := AIMSAcademicData{
		Email: body["email"].(string),
		Data:  receivedData,
	}

	saveData(AIMSdata)

	fmt.Fprint(w, "OK")
}
