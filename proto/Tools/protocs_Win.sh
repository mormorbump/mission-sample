#!/bin/bash

set -eu

PROTOC=$PWD/Tools/grpctools/tools/windows_x64/protoc.exe
GRPC_PLUGIN=$PWD/Tools/grpctools/tools/windows_x64/grpc_csharp_plugin.exe
UNITY_OUTPUT_PATH=../PoC/Assets/10_Scripts/Protos/

# Check existance
if [ ! -d $UNITY_OUTPUT_PATH ]; then
  # 存在しない場合は作成
  mkdir -p $UNITY_OUTPUT_PATH
fi


${PROTOC} \
  --csharp_out=$UNITY_OUTPUT_PATH\
  --grpc_out=$UNITY_OUTPUT_PATH \
  --plugin=protoc-gen-grpc=$GRPC_PLUGIN \
  -I ./src \
  ./src/*.proto
