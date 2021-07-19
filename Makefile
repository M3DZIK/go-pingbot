BINARY_NAME=pingbot.out

build:
	go build -o ${BINARY_NAME}

run:
	go run .

clean:
	go clean
	rm -rf pingbot* dist/
