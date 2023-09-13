##@ Other
.PHONY: commit-check
commit-check: ## commit pre check
	 pre-commit run --all-files

.PHONY: help
help: ## 展示帮助命令
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n"} \
			/^[a-zA-Z.]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 } \
			/^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) }' $(MAKEFILE_LIST)