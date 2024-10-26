#!/bin/bash

# Mac
brew update     # formula を更新

# Is already installed ?
IS_INSTALLED=`brew list | grep protobuf | wc -l`
echo $IS_INSTALLED
if [ ${IS_INSTALLED} -lt 1 ] ; then
	echo "Insatll protobuf"
	brew install protobuf # protobufをインストール
fi

brew upgrade protobuf # protobufをアップグレード

# VersionCheck
protoc --version