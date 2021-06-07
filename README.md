![build](https://github.com/crossphoton/iiitr-server/actions/workflows/go.yml/badge.svg)
# Server for the Services for the Institute

This is the package for the main Server of the services. Services are treated as extensions to the server.

The SERVER expects the following environment variables to be present
```
	- PORT 		- Port to serve on
	- DB_URL	- URL of a Postgres protocol supporting database
```
