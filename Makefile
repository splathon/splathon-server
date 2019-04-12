.PHONY: all
all: build

.PHONY: build
build:
	mkdir -p swagger
	swagger generate server -t ./swagger/ -A splathon -f=./splathon-api/swagger.yaml

.PHONY: install
install:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
