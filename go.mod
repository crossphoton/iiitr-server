module server

// +heroku goVersion go1.16
go 1.16

replace example.com/auth => ./auth/

replace example.com/studentdata => ./studentdata/

require (
	cloud.google.com/go v0.78.0 // indirect
	example.com/auth v0.0.0-00010101000000-000000000000
	example.com/studentdata v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	github.com/thanhpk/randstr v1.0.4 // indirect
)
