package auth

// Student type
type Student struct {
	Name         string
	Email        string `gorm:"primary_key;index"`
	Rollnumber   string `gorm:"index"`
	AccessToken  string
	RefreshToken string
	Picture      string
}
