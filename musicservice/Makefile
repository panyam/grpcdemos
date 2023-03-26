
# Some vars to detemrine go locations etc
GOROOT=$(which go)
GOPATH=$(HOME)/go
GOBIN=$(GOPATH)/bin

# Evaluates the abs path of the directory where this Makefile resides
SRC_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

# Where the protos exist
PROTO_DIR:=$(SRC_DIR)/protos

# where we want to generate server stubs, clients etc
OUT_DIR:=$(SRC_DIR)/gen/go/musicservice/v1

all: printenv goprotos gwprotos

goprotos:
	echo "Generating GO bindings"
	rm -Rf $(OUT_DIR) && mkdir -p $(OUT_DIR)
	protoc --go_out=$(OUT_DIR) --go_opt=paths=source_relative          	\
       --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative		\
       --proto_path=$(PROTO_DIR)/musicservice/v1/ 										\
       --proto_path=$(PROTO_DIR) 										\
      $(PROTO_DIR)/musicservice/v1/*.proto

gwprotos:
	echo "Generating gRPC Gateway bindings and OpenAPI spec"
	protoc -I . --grpc-gateway_out $(OUT_DIR) 													\
    --grpc-gateway_opt logtostderr=true 															\
    --grpc-gateway_opt paths=source_relative 													\
    --grpc-gateway_opt generate_unbound_methods=true 									\
    --proto_path=$(PROTO_DIR)/musicservice/v1/ 												\
    --proto_path=$(PROTO_DIR) 																				\
      $(PROTO_DIR)/musicservice/v1/*.proto

printenv:
	@echo MAKEFILE_LIST=$(MAKEFILE_LIST)
	@echo SRC_DIR=$(SRC_DIR)
	@echo PROTO_DIR=$(PROTO_DIR)
	@echo OUT_DIR=$(OUT_DIR)
	@echo GOROOT=$(GOROOT)
	@echo GOPATH=$(GOPATH)
	@echo GOBIN=$(GOBIN)
