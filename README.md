Golang sample codebase using Mux, Gorm and MySQL

Step to install:
- Clone this repository
- Create database `golang_mvc`
- Run application and database migration `go run main.go`

API endpoints:
- Create user `http://localhost:5000/users`
- Display users `http://localhost:5000/users`
- Display single user `http://localhost:5000/users/{id}`
- Update user `http://localhost:5000/users/{id}`
- Delete user `http://localhost:5000/users/{id}`

Run the unit test:
- Model test:
    -  `cd tests/modeltests`
    -  `go test -v`