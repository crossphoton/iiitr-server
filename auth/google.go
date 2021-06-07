package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/thanhpk/randstr"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	conf = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("DOMAIN") + "/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	} /* OAuth Config */

	state = randstr.String(10)
)

// googleOAuthRedirect redirects to the Google Page for account selection.
func googleOAuthRedirect(w http.ResponseWriter, r *http.Request) {
	// Checking if already looged in
	cookie, err := r.Cookie("token")
	if err == nil {
		_, err = GetClaims(cookie.Value)
	}
	if err != nil {
		// New login
		if r.URL.Query().Get("continue") != "" {
			conf.RedirectURL += "?continue=" + r.URL.Query().Get("continue")
		}
		url := conf.AuthCodeURL(state)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		return
	}

	if !Redirect(w, r) {
		w.Header().Set("Content-type", "text/plain")
		fmt.Fprintf(w, "Already logged in.\n")
	}
}

// Handles OAuth callback.
func googleOAuthCallback(w http.ResponseWriter, r *http.Request) {

	// Checking for state tampering
	if r.FormValue("state") != state {
		fmt.Println("state invalid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Getting tokens in exchange of cod
	token, err := conf.Exchange(context.Background(), r.FormValue("code"))

	if err != nil {
		fmt.Println("couldn't get token")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Getting profile information with acquired OAuth tokens
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	if err != nil {
		fmt.Println("couldn't get")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Deserialize response data
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("deserializing error")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Decode JSON data
	var data map[string]interface{}
	err = json.Unmarshal(content, &data)

	if err != nil {
		log.Fatal(err.Error())
	}

	jwtToken, err := (saveUser(data, token)) /* Save user and get JWT token */

	if err != nil && !Redirect(w, r) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Forming cookie
	dev, _ := strconv.ParseBool(os.Getenv("dev"))
	cookie := http.Cookie{Name: "token",
		Value:    jwtToken,
		Path:     "/",
		Expires:  time.Now().AddDate(10, 0, 0), /* 10 Year validity */
		Secure:   !dev,
		HttpOnly: true,
		//TODO: SameSite: ,
	}

	http.SetCookie(w, &cookie)
	if !Redirect(w, r) {
		fmt.Fprint(w, "Login successful!! You can close this tab now.")
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Query().Get("continue") == "" {
		return false
	}

	http.Redirect(w, r, r.URL.Query().Get("continue"), http.StatusTemporaryRedirect)
	return true
}
