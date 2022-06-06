.PHONY: build

build: briq-cli

briq-cli:	*.go briq/*.go render/*.go
	ENABLE_CGO=0 go build .

.PHONY: clean

clean:
	rm -f briq-cli
