run: build
	@./bin/api

build:
	@go build -o bin/api

clean:
	@rm -rf bin/*

test:
	@go test -v ./...