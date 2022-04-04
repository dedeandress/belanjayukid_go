# Go Sample Login Register

## Server
Heroku : https://go-sample-login-register.herokuapp.com

## Features

- Login (v1/auth/login)
- Register (v1/auth/register)
- User Detail (v1/me)

## Tech

Go Sample Login Register uses a number of open source projects to work properly:

- [Go] - Golang!
- [Postgres] - opensource database
- [Heroku] - Server.

## Installation

Go Sample Login Register requires [Go](https://golang.org/) v1.18 to run.

Build and start the server.

```sh
cd go_sample_login_register
export DATABASE_URL=CHANGEME
export DB_LOG_MODE=true #true or false
export DB_SSL_MODE=require #require or disable
export JWT_KEY=CHANGEME
go build main.go
./main.go
```

## Dependency

Go Sample Login Register is currently extended with the following Dependency.

| Dependency | URL |
| ------ | ------ |
| jwt-go | [github.com/dgrijalva/jwt-go/v4] |
| validator | [github.com/go-playground/validator/v10] |
| uuid | [github.com/google/uuid] |
| mux | [github.com/gorilla/mux] |
| schema | [github.com/gorilla/schema] |
| gorm | [github.com/jinzhu/gorm] |
| crypto | [golang.org/x/crypto] |

## License

MIT


[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [Go]: <https://golang.org>
   [Postgres]: <https://www.postgresql.org>
   [Heroku]: <http://heroku.com>
