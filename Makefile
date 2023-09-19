# Build options

PACKAGE=$(shell grep "^module" go.mod | awk '{print $$2}')

include scripts/make-rules/common.mk
include scripts/make-rules/tools.mk
include scripts/make-rules/golang.mk
include scripts/make-rules/format.mk
include scripts/make-rules/generate.mk
include scripts/make-rules/other.mk


##@ Bootstrap
.PHONY: init
init: format.install ## 初始化开发所需要的插件和工具

.DEFAULT_GOAL := help