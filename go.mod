module server

go 1.16

replace example.com/auth => ./auth/

replace example.com/studentdata => ./studentdata/

require (
	cloud.google.com/go v0.74.0 // indirect
	example.com/auth v0.0.0-00010101000000-000000000000
	example.com/studentdata v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
)
