package auth

import "example.com/studentdata"

func saveUser(s studentdata.Student) {
	db.Save(&s)
}
