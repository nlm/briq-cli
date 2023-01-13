.PHONY: build

build: briq-cli

briq-cli:	*.go briq/*.go commands/*.go config/*.go render/*.go
	ENABLE_CGO=0 go build .

.PHONY: clean

clean:
	rm -f briq-cli
