#!/bin/sh

echo $(protoc --go_out=. *.cgi.proto)
echo $(protoc --go_out=. *.data.proto)
echo $(protoc --go_out=. *.proto)
