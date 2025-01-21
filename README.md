# Project Go API Books

This is my project for learning Go.

In this repository, my preferred codebase style which I adopted from other repos, [1](https://github.com/fahminlb33/devoria1-wtc-backend) [2](https://github.com/dimaskiddo/codebase-go-rest). This project structure implements clean code and Domain-Driven Design (DDD).

## Libraries Used

This project uses several libraries to facilitate development and ensure code quality:

- **Fiber**: A web framework for building APIs in Go. It is used for routing and handling HTTP requests.
- **GORM**: An ORM library for Go. It is used for database interactions.
- **Viper**: A configuration management library. It is used for handling application configurations.
- **Zap**: A structured logger for Go. It is used for logging application events.
- **Testify**: A testing toolkit for Go. It is used for writing unit tests.
- **Mockery**: A mock code auto-generator for Go. It is used for generating mocks for testing.

## Project Structure

The project is structured as follows:

```
go-codebase/
│
├── cmd/main/           # Command line applications
│   ├── bootstrap.go    # Initializes and configures the application components
│   └── main.go         # Entry point of the application
│
├── pkg/                # Library code
│   ├── database/       # Database-related functionalities
│   └── ...             # Other packages
│
├── internal/           # Private application
│   ├── user/           # Domain business logic
│   └── ...             # Other domain business logic
│
├── api/                # API definitions and implementations
│   └── ...             # API-related code
│
├── configs/            # Configuration files
│   └── ...             # Configuration-related code
│
└── README.md           # Project documentation
```
