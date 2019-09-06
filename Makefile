default: debug

BIN_NAME := xsqlite
BIN_DIR := bin

.PHONY: directories

directories: ${BIN_DIR}

${BIN_DIR}:
	${MKDIR_P} ${BIN_DIR}

debug:
	@go build -o ${BIN_DIR}/${BIN_NAME} .
