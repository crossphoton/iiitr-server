package studentdata

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/crossphoton/iiitr-server/auth"
)

// updateData handles the request for AIMS data update
func updateAIMSData(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized")
		return
	}

	// Decoding data
	var AIMSdata AIMSAcademicData
	json.NewDecoder(r.Body).Decode(&AIMSdata)

	// Verifying user
	tokenStatus := auth.VerifyJWT(token.Value, AIMSdata.Email)
	if !tokenStatus {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized")
		return
	}
	AIMSdata.Timestamp = time.Now().Unix()

	saveData(AIMSdata)
	fmt.Fprint(w, "OK")
}

// getAIMSData handles AIMS data request
func getAIMSData(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized")
		return
	}

	claims, err := auth.GetClaims(token.Value)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized")
		return
	}

	searchResult := AIMSAcademicData{Email: claims["email"].(string)}
	db.First(&searchResult)
	jsonData, err := json.Marshal(searchResult)
	if err != nil {
		log.Println("error encoding to json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404: Not found")
		return
	}
	w.Write(jsonData)
}
