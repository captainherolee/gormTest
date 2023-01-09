#! /bin/bash

set -e


export SERVER_NAME="gormTest"
export CURRENT_DIR=$(pwd)
export SRC_DIR=$(pwd)
export GOMOD="$(pwd)/go.mod"
export VERSION="v0.1"

make clean
make build
