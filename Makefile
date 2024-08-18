.PHONY: help build gen-hello-grpc gen-hello-handler gen-hello

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

TARGET_MAX_CHAR_NUM=20

## Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

## Build the grpcgenie binary
build:
	go build -o grpcgenie cmd/grpcgenie/main.go

## Generate gRPC code for hello service
gen-hello-grpc:
	protoc --go_out=. --go_opt=paths=source_relative \
	 --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	 example/hello/hellopd/hello.proto

## Generate handler for hello service
gen-hello-handler:
	./grpcgenie --proto example/hello/hellopd/hello.proto \
	--output example/hello/hellohandler/hello_grpc_handler.go \
	--package hellohandler \
	--go-package-path github.com/pandakn/GrpcGenie/example/hello \
	--grpc-package hellopd 

## Generate both gRPC code and handler for hello service
gen-hello: gen-hello-grpc gen-hello-handler
