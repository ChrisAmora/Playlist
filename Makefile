init:
	go run github.com/99designs/gqlgen init
generate:
	go run github.com/99designs/gqlgen
run:
	go run main.go
test:
	go test -v