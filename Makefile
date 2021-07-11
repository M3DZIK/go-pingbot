BINARY_NAME=pingbot.out

build:
	go build -o ${BINARY_NAME}

all:
	./build.sh

run:
	go run .

deps:
	go get -v

clean:
	go clean
	rm pingbot* MD5* SHA256* VERSION
