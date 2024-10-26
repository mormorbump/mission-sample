#!/bin/bash

# Get gRPC Tools
curl -Lo grpctools.zip https://www.nuget.org/api/v2/package/Grpc.Tools/2.60.0
unzip grpctools.zip -d grpctools

# 権限
chmod +x grpctools/tools/macosx_x64/protoc
chmod +x grpctools/tools/macosx_x64/grpc_csharp_plugin