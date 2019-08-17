#!/bin/bash

GOOS=windows go build -o bin/curl-win.exe src/curl.go
GOOS=linux go build -o bin/curl-linux.exe src/curl.go
GOOS=darwin go build -o bin/curl-mac.exe src/curl.go