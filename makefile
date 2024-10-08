run:
	GOOS=darwin GOARCH=amd64 go build -o ./build/app ./cmd/main.go && ./build/app