# Auth

This application is an extension to [iiitr-services/Server](https://github.com/iiitr-services/Server)

Auth package implements authentication layer for user login

Currently supported methods
```
	- Google OAuth 			(for students)
```

This extension assumes these environment variables to be available (besides from main Server)
```
	- GOOGLE_CLIENT_ID			- For Google OAuth
	- GOOGLE_CLIENT_SECRET		- For Google OAuth
	- DOMAIN					- For cookie formation
	- JWT_SIGNING_KEY			- For JWT formation
```