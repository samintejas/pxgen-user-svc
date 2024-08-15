# Project x generic - user management service
A minimal microservice for usermanagement written in go

## How to use

mysql is required to run this application, table ddls are attached in the git

- Clone this repo
- Build the binary using go (v1.22 & above) `go build ./cmd/app/main.go -o user` , alternatively use `make build` OR
- use `go run ./cmd/app/main.go` to run the project without building alternatively use `make run`

### Currently available
- custom config
- custom logging
- http routing and middleware
- basic user apis
- login & auth

### under development
- validation
- security
- auditing
- admin apis
- roles
- root UI using HTMX
- api UI's using HTMX
