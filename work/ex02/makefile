GOCMD=go
GOBUILD=${GOCMD} build
GOCLEAN=${GOCMD} clean
GOTEST=${GOCMD} test
BINARY_NAME := bin/convert

all: build
build: 
	${GOBUILD} -o ${BINARY_NAME} -v

test: clean all
	${GOTEST} -v -cover
	cd ./convert; ${GOTEST} -v -cover

clean: 
	${GOCLEAN}
	rm -rf bin

.PHONY: test clean