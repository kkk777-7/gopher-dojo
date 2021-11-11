GOCMD=go
GOBUILD=${GOCMD} build
GOCLEAN=${GOCMD} clean
GOTEST=${GOCMD} test
BINARY_NAME := bin/omikuji

all: build
build: 
	${GOBUILD} -o ${BINARY_NAME} -v

clean: 
	${GOCLEAN}
	rm -rf bin

.PHONY: clean