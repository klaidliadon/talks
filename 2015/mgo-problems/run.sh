#!/bin/sh
export GOPATH=$PWD:$GOPATH
go generate node
go run src/main.go

