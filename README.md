# go-sample-api
Web api in golang with Gin farmwork:
```
GO-SAMPLE-API
├───.github                # GitHub-specific files
│   └───workflows          # CI/CD workflows
├───.vscode                # VSCode-specific files and configurations
├───cmd                    # Application entry points
│   └───api                # Main API entry point
│       └───routes         # Route configurations for the HTTP server
├───configs                # Configuration files
├───docs                   # Documentation files
├───internal               # Private application and library code
│   ├───controllers        # HTTP handlers
│   ├───core               # Core business logic
│   │   ├───entities       # Core business entities
│   │   ├───interfaces     # Interfaces for dependencies
│   │   │   ├───repositories  # Repository interfaces
│   │   │   └───services      # Service interfaces
│   │   └───usecases       # Business logic use cases
│   ├───dto                # Data transfer objects (request/response models)
│   └───infra              # Infrastructure implementations
│       ├───databases      # Database access implementations
│       ├───repositories   # Implementations of repository interfaces
│       └───services       # Implementations of service interfaces
├───pkg                    # Public library code
│   └───middlewares        # Custom middleware implementations
├───test                   # Test-related files
└───vendor                 # Third-party dependencies (managed by `go mod`)
```


## Update swagger document

```cmd
swag init -g cmd/api/main.go
```

## Important command

```
go mod int <package_name>
go get <package_url>
go mod tidy
go mod vendor
swag init -g cmd/api/main.go
```
## Debug it locally

Please install [delve](https://github.com/go-delve/delve/tree/master/Documentation/installation)

Click on debugger to run it.

## run it
```cmd
go run cmd/api/main.go
```
http://localhost:8080/swagger/index.html

![image](https://github.com/user-attachments/assets/87469707-8407-4076-9893-f429f04c4e7b)
