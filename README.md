# Mail for cause
This is a simple app built with vue and golang. this enables people to register their cause and then allow others to send mails directly from the app to concerning authorities. This project was inspired by [indiawantsbitcoin.org](https://indiawantsbitcoin.org).

## Project setup

### Environment Variables
This project required two environment variables:
- MONGO_URI: Mongodb URI
- SECRET:    A secret string used to encode and decode jwt

### Running the program
You can run the program by using the command `go run cmd/*go`

### Building the binaries
Running `go build cmd/*go` will build the binaries for the app.

_This app is open source. Feel free to contribute_