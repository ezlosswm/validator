test: 
	go test -v ./...

coverage: 
	go test -coverprofile=c.out

build: 
	@go build -o ./bin/main ./*.go

run: build 
	@./bin/main

clean: 
	@rm ./bin/main
