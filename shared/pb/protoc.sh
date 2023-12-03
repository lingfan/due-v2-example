#!/bin/bash

for file in `ls .`; do
    if [ -f ${file} ] && [ ${file##*.} == "proto" ]; then
      protoc --gofast_out=../../../ ${file}
    fi
done