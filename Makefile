.PHONY: all
all: build

build: propcheck.go
	go build -o propcheck_osx
	GOOS=linux GOARCH=amd64 go build -o propcheck_linux
