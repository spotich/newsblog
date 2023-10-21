build:
		go build -o bin/main.exe cmd/main.go

run:
		cd bin && ./main.exe

start: build run