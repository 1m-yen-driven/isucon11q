
SRC:=$(shell find . -type f -name '*.go' -not -path 'vendor/*')


.PHONY: all
all: build ;

.PHONY: build
build: isucondition

.PHONY: cross-compile
cross-compile: ${SRC}
	@GOOS=linux GOARCH=amd64 go build -o isucondition .

isucondition: ${SRC}
	@go build -o isucondition .
