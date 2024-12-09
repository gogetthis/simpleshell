.SILENT:

build: client server

client: bin/client
server: bin/server

bin/client:
	go build -o bin/client pkgs/client/main.go 

bin/server:
	go build -o bin/server pkgs/server/main.go

clean:
	rm -rf bin/**
	go mod tidy
	go clean -modcache

run: build
	./bin/server
