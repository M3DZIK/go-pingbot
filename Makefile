BINARY_NAME=pingbot.out

build:
	go build -o ${BINARY_NAME}

snapshot:
	goreleaser --snapshot --rm-dist

clean:
	go clean
	rm -rf pingbot* dist/
