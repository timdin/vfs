build:
	go generate ./...
	go build main.go

test:
	go generate ./...
	go test ./...