# GO Fiber Monolith Application

## Description
This is a full-stack monolith application using the GO Fiber framework for the backend and the HTML template engine for the frontend.
My project is researching the GO Fiber framework and how to use it to build a full-stack application.


## Packages Usage
- `go get github.com/jackc/pgx/v5`
- `go get -u github.com/gofiber/template/html/v2`
- `go get -u github.com/gofiber/fiber/v2`

## Installation
1. Clone the repository `git clone git@github.com:Marsudii/gofiber-monolith.git`
2. Run the command `go run cmd/main.go` to start the server
3. Open the browser and navigate to `http://localhost:3000`
4. You should see the home page of the application

## Structure Folder
- `cmd/` - contains the main.go file to start the server
- `internal/` - contains the application logic 
- `database/` - contains the database connection
- `handlers/` - contains the handler functions
- `services/` - contains the business logic'
- `repositories/` - contains the database queries
- `routes/` - contains the routes'
- `config/` - contains the configuration manually
- `models/` - contains the data models
- `views/` - contains the HTML templates


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.


## Author
- [Marsudi](https://github.com/Marsudii)