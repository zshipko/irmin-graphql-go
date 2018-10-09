#!/bin/sh

temp=`mktemp -d`
irmin graphql --root $temp --port 8080 > /dev/null &
PID=$!
go test -v
kill $PID
rm -r $temp
