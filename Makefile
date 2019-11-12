# Windows... :/
export GOOS=windows

proj := $(shell basename $(shell pwd).exe)


default: all

all: $(proj)


check:
	errcheck .

test:
	go test

format:
	gofmt -s=true -w *.go

clean:
	rm -f *~ */*~ .*~ $(proj)
	go clean

$(proj): Makefile *.go
	go vet
	go build -ldflags="-w -s"

.PHONY: all check test format clean
