#!/bin/sh -e

INC=$(go list -f '{{ .Dir }}' -m github.com/unistack-org/micro-proto)
ARGS="-I${INC}"

protoc $ARGS -Iproto --openapiv2_out=disable_default_errors=true,allow_merge=true:./proto/ --go_out=paths=source_relative:./proto/ --micro_out=components="micro|http",debug=true,paths=source_relative:./proto/ proto/*.proto