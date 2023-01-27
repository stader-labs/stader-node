#!/bin/bash

export CGO_ENABLED=0
cd /stader-node/stader-cli

# Build x64 version
GOOS=linux GOARCH=amd64 go build -o stader-cli-linux-amd64 stader-cli.go
GOOS=darwin GOARCH=amd64 go build -o stader-cli-darwin-amd64 stader-cli.go

# Build the arm64 version
GOOS=linux GOARCH=arm64 go build -o stader-cli-linux-arm64 stader-cli.go
GOOS=darwin GOARCH=arm64 go build -o stader-cli-darwin-arm64 stader-cli.go
