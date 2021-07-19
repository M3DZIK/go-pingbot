BINARY_NAME=pingbot.out

build:
	go build -o ${BINARY_NAME}

run:
	go run .

deps:
	go get -v

clean:
	go clean
	rm pingbot* dist/
