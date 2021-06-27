.PHONY: build
build:
	CGO_ENABLED=0 go build -o bin/generator cmd/generator/main.go
	CGO_ENABLED=0 go build -o bin/sorter cmd/sorter/main.go