# go-sample-api
Web api in golang with Gin farmwork


## Update swagger document

```cmd
swag init -g cmd/api/main.go
```

## Important command

```
go mod int <package_name>
go get <package_url>
go mod vendor
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