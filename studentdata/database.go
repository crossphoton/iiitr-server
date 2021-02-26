package studentdata

func saveData(d AIMSAcademicData) {
	db.Save(&d)
}
