#!/bin/bash
until [ -d ../../../src/scheduling ]; do
echo "Retrying directory..."
sleep 1
done

if [ $TEST = '1' ]
then
    cd /go/src/calendar
    go get github.com/2tvenom/go-test-teamcity
    go test -v ./... | go-test-teamcity
    sleep 10
    for i in {1..500}
    do
        go test -v ./... | go-test-teamcity
        sleep 5
    done
else
    go get ./
    go build
    realize start --run
fi
