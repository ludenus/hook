#!/bin/bash
set -x
set -e

pwd 
ls -pilaF
git log -1
go version

go get ./...

GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -v github.com/ludenus/hook 

chmod 777 ./hook
