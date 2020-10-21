build:
	if [ -d "./bin" ]; then rm -r bin; fi
	mkdir bin
	go build -o bin/nospin cmd/nospin.go

run:
	go run main.go

release:
	if [ -d "./bin" ]; then rm -r bin; fi
	if [ -d "./release" ]; then rm -r release; fi
	mkdir release
	mkdir bin
	go build -o bin/nospin cmd/nospin.go
	cp bin/nospin release
	cp README.md release
	zip -9 -q -r release.zip release
	rm -r release
