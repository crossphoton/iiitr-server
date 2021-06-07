package studentdata

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

// Check database connection
func CheckDB() {
	err = db.DB().Ping()
	if err != nil {
		log.Println("Reconnecting to DB")
	}

	addr := os.Getenv("DB_URL")
	db, err = gorm.Open("postgres", addr)
	db.AutoMigrate(AIMSAcademicData{})

	if err != nil {
		log.Fatal(err)
		log.Fatal("DB Error")
	}
}
