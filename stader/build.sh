#!/bin/bash

export CGO_ENABLED=1
cd /stader-node/stader

# GOOS=darwin GOARCH=amd64 go build -o stader-cli-darwin-amd64 stader-cli.go

#
## Build x64 version
CGO_CFLAGS="-O -D__BLST_PORTABLE__" GOARCH=amd64 GOOS=linux go build -o stader-daemon-linux-amd64 stader.go
#
## Build the arm64 version
CC=aarch64-linux-gnu-gcc CXX=aarch64-linux-gnu-cpp CGO_CFLAGS="-O -D__BLST_PORTABLE__" GOARCH=arm64 GOOS=linux go build -o stader-daemon-linux-arm64 stader.go