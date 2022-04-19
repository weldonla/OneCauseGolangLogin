# OneCauseGolangLogin
This project was generated with Go cli.

## Development Server
Run `go run .` to start server, running on http://127.0.0.1:8000

## MySql Server
For this example to work, you will need to make a new schema in MySql called 'userdb' (or just make the appropriate changes in the databases > user.go file). You will 
likely need to go to that user.go file anyway to replace 'root:root' with whatever your username and password is for your local MySql server.
## Sources
This Golang project is adapted from the [following project](https://codesource.io/how-to-setup-golang-authentication-with-jwt-token/).