package studentdata

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Student type
type Student struct {
	Name         string
	Email        string `gorm:"primary_key"`
	Rollnumber   string
	AccessToken  string
	RefreshToken string
	Picture      string
}

// AIMSAcademicData type
type AIMSAcademicData struct {
	Email string `gorm:"primary_key"`
	Data  string
}
