.PHONY: build check fix test

BINARY=logcheck
TESTDIR=./testdata/src/a/...

build:
	go build -o $(BINARY) ./cmd/logcheck
check: build
	go vet -vettool=./logcheck $(TESTDIR)
fix: build
	go vet -vettool=./logcheck --fix $(TESTDIR)
test:
	go test ./...
