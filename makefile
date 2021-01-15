build:
	if [ -d "./bin" ]; then rm -r bin; fi
	mkdir bin
	if [ -d "./bin" ]; then rm -r bin; fi
	go build -o bin/qt cmd/quote.go

run:
	go run main.go

install:
	if [ -d "./bin" ]; then rm -r bin; fi
	mkdir bin
	go build -o bin/qt cmd/quote.go
	sudo mv bin/qt /usr/bin
	# sudo mv web /usr/bin
