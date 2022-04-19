# OneCauseGolangLogin
This project was generated with Go cli.

## Development Server
Run `go run .` to start server, running on http://127.0.0.1:8000

## MySql Server
For this example to work, you will need to make a new schema in MySql called 'userdb' (or just make the appropriate changes in the databases > user.go file). You will 
likely need to go to that user.go file anyway to replace 'root:root' with whatever your username and password is for your local MySql server. You will also need to insert
the necessary user into the database, which should have the following information:
* "id": 0
* "name": "Rick"
* "email": "c137@onecause.com"
* "password": "JDJhJDEyJDRXMGlFekFMR0huTi9zeDlBc2FDZGVHNWhpZnJIM3VPOVhuWFR5V01OTERoODBkMUU0b1Zx"
* "user_name": "c137@onecause.com"

Realistically though, the only actual important information is the email and password.

## Sources
This Golang project is adapted from the [following walkthrough](https://codesource.io/how-to-setup-golang-authentication-with-jwt-token/).