package studentdata

// AIMSAcademicData type
type AIMSAcademicData struct {
	Email string `gorm:"primary_key"`
	Data  string
}
