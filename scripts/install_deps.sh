#!/bin/bash

set -ex
OS=$(uname -s)

export GO111MODULE=on

go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4

if [[ "$OS" == "Darwin" ]]; then
	if command -v brew > /dev/null 2>&1; then
		brew install jq
	fi
else
  echo "Please install 'jq': https://stedolan.github.io/jq/download/"
fi