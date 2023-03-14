#!/bin/sh
dir=$(ls |awk '/^*\.rpc/')

for file in $dir
do
protoc  --go_out="../gameMsg" --plugin=grpc:. $file
echo $file
done
