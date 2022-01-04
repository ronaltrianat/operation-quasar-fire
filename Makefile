PKG_LIST := $(shell go list ./... | grep -v /vendor/)

default: build

all: test build

tidy:  ## Execute tidy comand
	@go mod tidy

build: tidy ## Build the binary file
	go build -i -v $(PKG_LIST)

docker-build:
	docker build -t meli-tl-challenge .

docker-run:
	docker run -d -p 9000:9000 meli-tl-challenge

fmt: ## Formmat src code files
	@go fmt ${PKG_LIST}

test: ## Execute test
	@echo "go test ${PKG_LIST}"
	go test ./tests || exit 1

race: ## Run data race detector
	@go test -race -short ${PKG_LIST}

mocks:
	@mockgen -source=src/core/ports/message_port.go -destination=tests/mocks/mockups/message_port.go
	@mockgen -source=src/core/ports/trilateration_port.go -destination=tests/mocks/mockups/trilateration_port.go
	@mockgen -source=src/core/ports/repository_port.go -destination=tests/mocks/mockups/repository_port.go