compile:
	GOOS=darwin GOARCH=amd64 go build -o ./build/app_amd64 ./cmd/main.go
run.listen:
	./build/app_amd64 -address 127.0.0.1:3001
run.send:
	./build/app_amd64 -send -address 127.0.0.1:3001 -file ./from/test_1.txt
