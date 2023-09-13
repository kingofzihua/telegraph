
VERSION_PACKAGE = "$(PACKAGE)/pkg/version"

GO_LDFLAGS += -X $(VERSION_PACKAGE).GitVersion=$(VERSION) \
	-X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PACKAGE).GitTreeState=$(GIT_TREE_STATE) \
	-X $(VERSION_PACKAGE).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

##@ Golang
.PHONY: go.build
go.build:## go build
	mkdir -p bin/ && go build -ldflags "$(GO_LDFLAGS)" -o ./bin/ ./...