# Machship Exam

This app serves as an exam for my Golang developer application.

## Steps to Run the App

1. Clone the repository.
2. Run the program using the following command in your command prompt or terminal:
```shell
go run cmd/api/main.go
```

## Sample Command Prompt Request

You can use `curl` to make a sample request to the application. Here's an example:

```shell
curl --location --request GET 'http://127.0.0.1:8080/api/v1/retrieveUsers' \
--header 'Content-Type: application/json' \
--data '{
 "usernames": ["TheAlgorithms", "nonExistingId12345", "torvalds", "bradtraversy", "torvalds", "michaelliao", "mamontes1788"]
}'
```
3. Run unit tests (optional)
- Handler Level
```shell
go test ./internal/network/http/handler
```
- Service Level
```shell
go test ./internal/core/service
```