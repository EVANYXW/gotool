#!/bin/sh
dir=$(ls |awk '/^*.proto/')

for file in $dir
do
protoc  --go_out="./Msg" --plugin=grpc:. $file
done
