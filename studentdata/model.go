package studentdata

import "gorm.io/datatypes"

// AIMSAcademicData type as in current system
type AIMSAcademicData struct {
	Email     string `json:"email" gorm:"primary_key;index"`
	Data      datatypes.JSON
	Timestamp int64 `json:"timestamp"`
}
