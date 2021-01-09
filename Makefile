all: build

build: statik main

main:
	@CGO_ENABLED=0 go build -o ./build/app main.go

statik:
	@statik -src ./dist

install_deps:
	@go get github.com/rakyll/statik

.PHONY: build statik main install_deps
