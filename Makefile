ENV_FILE=.env

install:
	@echo "-> Creating .env file"
	@if test -f ${ENV_FILE}; then echo ".env file exists already"; exit 1; fi
	@mkdir -p $(shell dirname ${ENV_FILE}) && touch ${ENV_FILE}
	@echo ROOT=\"${shell dirname $$(realpath -s ${ENV_FILE})}\" > ${ENV_FILE}

build:
	@echo "-> Building Library"
	@go build -o compiled/libmatch.so -buildmode=c-shared