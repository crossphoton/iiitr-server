package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"example.com/studentdata"
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
	}

	state = os.Getenv("OAUTH_STATE") // TODO: Randomize it
)

func googleOAuthRedirect(w http.ResponseWriter, r *http.Request) {
	url := conf.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func googleOAuthCallback(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("state") != state {
		fmt.Println("state invalid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect) //  TODO: Improve
		return
	}

	token, err := conf.Exchange(oauth2.NoContext, r.FormValue("code"))

	if err != nil {
		fmt.Println("couldn't get token")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect) //  TODO: Improve
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	if err != nil {
		fmt.Println("couldn't get")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect) //  TODO: Improve
		return
	}

	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("deserializing error")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect) //  TODO: Improve
		return
	}

	var data map[string]interface{}

	err = json.Unmarshal(content, &data)

	if err != nil {
		log.Fatal(err.Error())
	}

	stu := studentdata.Student{
		Name:         data["name"].(string),
		Email:        data["email"].(string),
		Rollnumber:   strings.Split(data["email"].(string), "@")[0],
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Picture:      data["picture"].(string),
	}

	saveUser(stu)

	fmt.Fprintf(w, "Response: %s", content)

}
