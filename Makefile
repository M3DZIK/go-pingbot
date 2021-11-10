BINARY_NAME := pingbot.out

# executables
GO         := go
GORELEASER := goreleaser

# build flags
BUILD_FLAGS :=

build:
	$(GO) mod tidy
	$(GO) build $(BUILD_FLAGS) -o $(BINARY_NAME)

snapshot:
	$(GORELEASER) --snapshot --rm-dist

clean:
	$(GO) clean
	rm -rf pingbot* dist/
