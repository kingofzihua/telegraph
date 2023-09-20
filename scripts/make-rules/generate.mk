##@ Generate

.PHONY: gen.protoc
gen.protoc: ## 编译 protobuf 文件.
	@echo "===========> Generate protobuf files start"
	@protoc                                            \
    		--proto_path=$(PROTO_DIR)                          \
    		--proto_path=$(ROOT_DIR)/third_party             \
    		--go_out=paths=source_relative:$(PROTO_DIR)        \
    		--go-grpc_out=paths=source_relative:$(PROTO_DIR)   \
    		$(shell find $(PROTO_DIR) -name *.proto)
	@echo "===========> Generate protobuf files success"

.PHONY: gen.deps
gen.deps: tools.verify ## 安装依赖，例如：生成需要的代码等.
	@go generate $(ROOT_DIR)/...


.PHONY: gen.models
gen.models: ## gorm generate models
	go run ./tools/gormgen/main.go -c .gormgen.yaml -outPath "./internal/apicore/data/query"