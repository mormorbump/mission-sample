#!/bin/bash

set -eu

PROTOC=$PWD/Tools/grpctools/tools/macosx_x64/protoc
GRPC_PLUGIN=$PWD/Tools/grpctools/tools/macosx_x64/grpc_csharp_plugin
UNITY_OUTPUT_PATH=../PoC/Assets/10_Scripts/Protos/

# Check existance
if [ ! -d $UNITY_OUTPUT_PATH ]; then
  # 存在しない場合は作成
  mkdir -p $UNITY_OUTPUT_PATH
fi


protoc \
  --csharp_out=$UNITY_OUTPUT_PATH\
  --grpc_out=$UNITY_OUTPUT_PATH \
  --plugin=protoc-gen-grpc=$GRPC_PLUGIN \
  -I ./ \
  ./*.proto
