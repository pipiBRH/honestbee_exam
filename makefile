all: clean build rmlog

clean:
	go clean src/main.go
build: 
	go build  -o honestbee src/main.go

rmlog:
	rm ./log/*

linux: clean rmlog
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o linux_amazon_spider	