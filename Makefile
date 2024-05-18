install:
	go mod tidy
	go mod download
	CGO_ENABLED=0 go build -o $(HOME)/go/bin/superkey

