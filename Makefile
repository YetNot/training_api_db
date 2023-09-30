.PHONY: build

build:
	go build -v -buildvcs=false ./cmd/api/
run:
	./api.exe
	
.DEFAULT_GOAL:= build