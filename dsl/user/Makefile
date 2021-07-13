BINARY_NAME=user

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} -v main.go

run:
	go run main.go