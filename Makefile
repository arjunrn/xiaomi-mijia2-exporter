.PHONY: lint build

all: build

lint:
	golangci-lint run -c .golangci.yaml 

build:
	go build -o bin/xiaomi-mijia2-exporter main.go
