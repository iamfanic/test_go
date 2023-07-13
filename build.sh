#!/bin/bash -l

source /etc/profile
go env
go mod tidy
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go