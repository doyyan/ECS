go mod tidy
go build
go test ./... -v -cover
./shoppingtrip
