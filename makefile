build:
	if [ -d "./bin" ]; then rm -r bin; fi
	mkdir bin
	if [ -d "./bin" ]; then rm -r bin; fi
	go build -o bin/nospin cmd/nospin.go

run:
	go run main.go

install:
	if [ -d "./bin" ]; then rm -r bin; fi
	mkdir bin
	go build -o bin/nospin cmd/nospin.go
	sudo mv bin/nospin /usr/bin