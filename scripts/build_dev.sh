#!/usr/bin/env bash

if [[ "$0" != *scripts/*.sh ]]
then
	echo "Please, execute from project's root directory"
	exit 1
fi

source scripts/build_common.inc.sh

mkdir -p bin
rm -f bin/$EXE_NAME*

GOOS=darwin GOARCH=amd64 go build -ldflags "${FLAGS}" -o tmp/retroserver retroserver.go
if [[ $? -ne 0 ]]
then
    echo "Compilation error!"
    exit 1
fi
