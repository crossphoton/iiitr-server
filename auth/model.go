package auth

// Student type
type Student struct {
	Name         string
	Email        string `gorm:"primary_key"`
	Rollnumber   string
	AccessToken  string
	RefreshToken string
	Picture      string
}
