default: debug

BIN_NAME := xsqlite
BIN_DIR := bin
COMMIT_HASH := $(shell git rev-parse --short HEAD)
VERSION_NUMBER := v0.1.0
VERSION_FLAG := -X xk0nsid/xsqlite/cmd.versionNumber=${VERSION_NUMBER}
COMMIT_FLAG := -X xk0nsid/xsqlite/cmd.gitRevNumber=${COMMIT_HASH}
LDFLAGS := "${VERSION_FLAG} ${COMMIT_FLAG}"

.PHONY: directories

directories: ${BIN_DIR}

${BIN_DIR}:
	${MKDIR_P} ${BIN_DIR}

debug:
	@go build -ldflags ${LDFLAGS} -o ${BIN_DIR}/${BIN_NAME} .
