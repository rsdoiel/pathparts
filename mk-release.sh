#!/bin/bash
#
# Make releases for Linux/amd64, Linux/ARM7 (Raspberry Pi), Windows, and Mac OX X (darwin)
#
PROGNAME=pathparts

env GOOS=linux GOARCH=amd64 go build -o dist/linux-amd64/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
env GOOS=linux GOARCH=arm GOARM=6 go build -o dist/raspberrypi-arm6/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/raspberrypi-arm7/$PROGNAME cmds/$PROGNAME/$PROGNAME.go
env GOOS=windows GOARCH=amd64 go build -o dist/windows/$PROGNAME.exe cmds/$PROGNAME/$PROGNAME.go
env GOOS=darwin	GOARCH=amd64 go build -o dist/maxosx/$PROGNAME cmds/$PROGNAME/$PROGNAME.go

