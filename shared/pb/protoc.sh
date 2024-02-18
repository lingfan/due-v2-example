#!/bin/bash

# for file in `ls .`; do
#     if [ -f ${file} ] && [ ${file##*.} == "proto" ]; then
#       protoc --gofast_out=../../../ ${file}
#     fi
# done

# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

for file in `ls .`; do
    if [ -f ${file} ] && [ ${file##*.} == "proto" ]; then
      protoc --go_out=./out ${file}
    fi
done

cp -r ./out/due-v2-example/shared/pb/* ./
rm -rf ./out