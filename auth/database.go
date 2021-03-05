package auth

import (
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/oauth2"
)

var db *gorm.DB
var err error

func saveUser(data map[string]interface{}, token *oauth2.Token) (string, error) {

	stu := Student{
		Name:         data["name"].(string),
		Email:        data["email"].(string),
		Rollnumber:   strings.Split(data["email"].(string), "@")[0],
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Picture:      data["picture"].(string),
	}

	checkDB()

	db.Save(&stu)

	return generateJWTToken(stu)
}

func checkDB() {

	err = db.DB().Ping()
	if err != nil {
		log.Println("Reconnecting to DB")
		addr := os.Getenv("DB_URL")
		db, err = gorm.Open("postgres", addr)

		if err != nil {
			log.Fatal(err)
			log.Fatal("DB Error")
		}
	}
}
