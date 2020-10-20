build:
	if [ -d "./bin" ]; then rm -r bin; fi
	mkdir bin
	go build -o bin/nospin cmd/nospin.go

run:
	go run main.go