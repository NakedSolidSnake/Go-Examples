hello:
	echo "Hello"

build:
	go build -o bin/main main.go

run:
	go run main.go 

compile:
	GOOS=linux GOARCH=386 go build -o bin/main-386 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/main-amd64 main.go
	GOOS=linux GOARCH=arm go build -o bin/main-arm main.go

