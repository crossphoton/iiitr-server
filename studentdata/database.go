package studentdata

// Save received data from AIMS
func saveData(d AIMSAcademicData) {
	CheckDB()
	db.Save(&d)
}
