#
# Simple Makefile
#

build:
	go build -o bin/pathparts cmds/pathparts/pathparts.go 

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

install:
	env GOBIN=$(HOME)/bin go install cmds/pathparts/pathparts.go

release:
	./mk-release.sh
