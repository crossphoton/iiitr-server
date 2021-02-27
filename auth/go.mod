module auth

go 1.16

replace example.com/studentdata => ../studentdata/

require (
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	golang.org/x/oauth2 v0.0.0-20210220000619-9bb904979d93
)
