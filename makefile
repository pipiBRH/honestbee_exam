all: clean build

clean:
	go clean
download:
	go mod download	
build: 
	go build  -o ./bin/main
linux: clean download
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main