# Gin RESTful

This project demonstrates the creation of a basic RESTful API using the Golang programming language with the Gin web framework.

## Project Structure

The project file and folder structure is as follows:

    - configs
    - controllers
      - todo_controller.go
    - db
      - db.go
      - migration.go
    - interfaces
    - middlewares
      - auth_middleware.go
      - cors_middleware.go
      - rate_limit_middleware.go
    - models
      - base.go
      - todo.go
    - repositories
      - abstract_repository.go
      - todo_repository.go
    - routes
      - auth.go
      - routes.go
      - todos.go
    - services
      - abstract_service.go
      - todo_service.go
    - utils
      - utils.go
    - .env
    - .env.example
    - Dockerfile
    - go.mod
    - go.sum
    - main.go

  
## Key Concepts

- **Gin:** A Golang web framework used for creating web applications and microservices. For more information: [Gin Framework](https://github.com/gin-gonic/gin)
- **GORM:** Used for Object Relational Mapping (ORM) in Golang, managing database operations. For more information: [GORM ORM](https://gorm.io/)
- **Docker:** Used to containerize the application. For more information: [Docker](https://www.docker.com/)

## Examining Main Files

- **main.go:** The starting point of the application. Sets up the router, manages connections, and starts the server.
- **db/db.go:** Initializes the connection to the MySQL database using GORM.
- **middlewares/auth_middleware.go:** Middleware used for authentication.
- **routes/auth.go:** File defining endpoints related to authentication.
- **models/todo.go:** Model representing Todo items.

## Database Initialization

To initialize your database and create tables, run the following command:

```bash
go run main.go migrate
```

## Running the Application
You can run the application with the following command:
```bash
go run main.go
```

The application will run on localhost:8080 by default.

## Contributing
- Fork this repository ("Fork" button at the top).
- Create a new branch (git checkout -b feature/fooBar).
- Commit your changes (git commit -am 'Add some fooBar').
- Push to your branch (git push origin feature/fooBar).
- Open a Pull Request.

## License
This project is licensed under the MIT License. See [LICENSE](LICENSE) for more information.
