##@ Tools
# ==============================================================================
# 工具相关的 Makefile
#

TOOLS ?= golangci-lint goimports gotests mockgen protoc-gen-go swagger addlicense migrate

MIGRATE_VERSION="v4.15.2"


.PHONY: tools.verify
tools.verify: $(addprefix tools.verify., $(TOOLS)) ## 验证工具

.PHONY: tools.install
tools.install: $(addprefix tools.install., $(TOOLS)) ## 安装工具

.PHONY: tools.install.%
tools.install.%:
	@echo "===========> Installing $*"
	@$(MAKE) install.$*

.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: install.golangci-lint
install.golangci-lint:
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
	@golangci-lint completion bash > $(HOME)/.golangci-lint.bash
	@if ! grep -q .golangci-lint.bash $(HOME)/.bashrc; then echo "source \$$HOME/.golangci-lint.bash" >> $(HOME)/.bashrc; fi

.PHONY: install.goimports
install.goimports:
	@$(GO) install golang.org/x/tools/cmd/goimports@latest

.PHONY: install.gotests
install.gotests:
	@$(GO) install github.com/cweill/gotests/gotests@latest

.PHONY: install.mockgen
install.mockgen: ## install mockgen
	@$(GO) install github.com/golang/mock/mockgen@latest

.PHONY: install.protoc-gen-go
install.protoc-gen-go:
	@$(GO) install github.com/golang/protobuf/protoc-gen-go@latest

.PHONY: install.swagger
install.swagger:
	@$(GO) install github.com/go-swagger/go-swagger/cmd/swagger@latest

.PHONY: install.addlicense
install.addlicense:
	@$(GO) install github.com/marmotedu/addlicense@latest

.PHONY: install.migrate
install.migrate: ## install migrate
	@$(GO) install -ldflags='-X main.Version=$(MIGRATE_VERSION) -extldflags "-static"' -tags '$(DATABASE_DRIVER)' github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION)
	@echo "migrate@$(MIGRATE_VERSION) install success"
