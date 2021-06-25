#!/bin/sh

echo $(protoc --go_out=. user.cgi.proto)
echo $(protoc --go_out=. user.data.proto)
