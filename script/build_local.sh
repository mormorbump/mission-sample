#!/usr/bin/env bash

set -eu

docker build \
  --platform linux/arm64 \
  -t graffity_centray_api_local \
  -f Dockerfile .

