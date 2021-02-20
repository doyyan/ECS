go mod tidy
go build
go test ./... -v -cover 
go tool cover -func=coverage.out
go build shoppingtrip.go 
./shoppingtrip
