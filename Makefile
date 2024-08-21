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
	@go build -o go-wordle ./main
	@echo "Build completed: go-wordle"

run:
	@./go-wordle

clean:
	@echo "Cleaning.."
	@rm -f go-wordle
	@echo "Cleaned"