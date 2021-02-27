package studentdata

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func saveData(d AIMSAcademicData) {
	err = db.DB().Ping()
	if err != nil {
		dbInit()
		log.Println("Reconnecting to DB")
	}

	db.Save(&d)
}

func dbInit() {
	addr := os.Getenv("DB_URL")
	db, err = gorm.Open("postgres", addr)

	if err != nil {
		log.Fatal(err)
		log.Fatal("DB Error")
	}
}
