.PHONY: build

build:
	go build -v ./cmd/api/
hello:
	./api.exe
	
.DEFAULT_GOAL:= build