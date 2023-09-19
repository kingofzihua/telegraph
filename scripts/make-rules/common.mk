COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
# 项目根目录
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/../../ && pwd -P))

PACKAGE=$(shell grep "^module" go.mod | awk '{print $$2}')

# Protobuf 文件存放路径
PROTO_DIR=$(ROOT_DIR)/proto


# ==============================================================================
# 定义版本相关变量

ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif
# Check if the tree is dirty.  default to dirty
GIT_TREE_STATE:="dirty"
ifeq (, $(shell git status --porcelain 2>/dev/null))
	GIT_TREE_STATE="clean"
endif
GIT_COMMIT:=$(shell git rev-parse HEAD)


# ==============================================================================
# 定义golang相关变量
GO := go