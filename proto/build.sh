#!/bin/sh

function common() {
  echo $(protoc --go_out=. common.proto)
  return $?
}

function user() {
  echo $(protoc --go_out=. user/*.proto)
  return $?
}

case $1 in
'user')
  user
  ;;

*)
  common
  user
  ;;
esac

exit $?
