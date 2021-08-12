ENV_FILE=.env

help:
	@echo "Supported Commands"
	@echo "build :  Build a C-Shared Library from source code"
build:
	@echo "-> Building Library"
	@go build -o compiled/libmatch.so -buildmode=c-shared