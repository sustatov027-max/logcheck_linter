build:
	go build -o logcheck ./cmd/logcheck
check: build
	go vet -vettool=./logcheck ./testdata/src/a/a.go
fix: build
	go vet -vettool=./logcheck --fix ./testdata/src/a/a.go
