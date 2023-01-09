CURRENT_BIN_DIR=${SRC_DIR}/bin
BIN_DIR := ${SRC_DIR}/bin
VERSION := ${VERSION}

default: build

build: bins server

clean:
	rm -rf ${CURRENT_BIN_DIR}

bins:
	mkdir ${BIN_DIR}

server:
	go build -a -v -o  ${BIN_DIR}/${SERVER_NAME} ${SRC_DIR}/*.go

.PHONY: build clean server bins

