run:
	go run ./cmd/hotreload/main.go

#windows
buildw:

	rm -rf ./bin/win
	go build -o ./bin/win/hotreload.exe ./cmd/hotreload/main.go

#linux
buildl:

	rm -rf ./bin/linux
	GOOS=linux GOARCH=amd64 go build -o ./bin/linux/hotreload ./cmd/hotreload/main.go
