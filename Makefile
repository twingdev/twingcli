

TARGET_DIR = "./cli/"
BINARY_NAME = "${TARGET_DIR}twing-cli"

all: build

.PHONY: all

compile_all:
		GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go && GOOS=linux go build -o ${BINARY_NAME}-linux main.go && GOOS=windows go build -o ${BINARY_NAME}-windows main.go

compile_win:
	GOOS=windows go build -o ${BINARY_NAME}-windows main.go

compile_darwin:
	GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go


build: clean compile_all
	mkdir -p cli
	@echo "compiling twing cli to: ${TARGET_DIR}"

clean:
	rm -rf cli

protoc:
	protoc --go_out=. server/pb/*.proto
