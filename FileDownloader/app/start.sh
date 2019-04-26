#!/bin/bash

if [[ $TEST = '1' ]]
then
    cd /go/src/downloader
    go get github.com/2tvenom/go-test-teamcity
    go test -v ./... | go-test-teamcity
else
    go get ./
    go build
    realize start --run
fi