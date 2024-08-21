all: lint test build run

get-tools:
	@go install golang.org/x/lint/golint@latest

lint:
	@echo "Started linting.."
	@golint ./...

test:
	@echo "Started testing.."
	@go test ./words
	@go test ./verifier
	@go test ./game

build:
	@echo "Buildling.."
	@go build -o go-wordle main.go
	@GOARCH=amd64 GOOS=darwin go build -o go-wordle-darwin main.go
	@GOARCH=amd64 GOOS=linux go build -o go-wordle-linux main.go
	@GOARCH=amd64 GOOS=windows go build -o go-wordle-windows main.go
	@echo "Build completed: go-wordle"

run:
	@./go-wordle-darwin

clean:
	@echo "Cleaning.."
	@rm -f go-wordle-darwin
	@rm -f go-wordle-linux
	@rm -f go-wordle-windows
	@echo "Cleaned"