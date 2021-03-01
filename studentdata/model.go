package studentdata

// AIMSAcademicData type
type AIMSAcademicData struct {
	Email     string `gorm:"primary_key;index"`
	Data      string
	Timestamp int64 `gorm:"autoCreateTime"`
}
