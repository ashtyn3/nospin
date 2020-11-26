build:
	if [ -d "./bin" ]; then rm -r bin; fi
	mkdir bin
	if [ -d "./bin" ]; then rm -r bin; fi
	go build -o bin/qoute cmd/qoute.go

run:
	go run main.go

install:
	if [ -d "./bin" ]; then rm -r bin; fi
	mkdir bin
	go build -o bin/qoute cmd/qoute.go
	sudo mv bin/qoute /usr/bin