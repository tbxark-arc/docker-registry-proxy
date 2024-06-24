BIN_NAME=docker-registry-mirror
BUILD_DIR=./build
BUILD=$(shell git rev-parse --short HEAD)@$(shell date +%s)
CURRENT_OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
CURRENT_ARCH := $(shell uname -m | tr '[:upper:]' '[:lower:]')
LD_FLAGS=-ldflags "-X main.BuildVersion=$(BUILD)"
GO_BUILD=CGO_ENABLED=0 go build $(LD_FLAGS)

.PHONY: build
build:
	$(GO_BUILD) -o ./build/$(BIN_NAME)_$(CURRENT_OS)_$(CURRENT_ARCH)/ ./...

.PHONY: buildLinuxX86
buildLinuxX86:
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(BUILD_DIR)/$(BIN_NAME)_linux_x86/ ./...

.PHONY: deploy
deploy: buildLinuxX86
	@echo "Deploying to tecent"
	@scp $(BUILD_DIR)/$(BIN_NAME)_linux_x86/$(BIN_NAME) tecent:~/$(BIN_NAME)
	@ssh tecent "sudo systemctl stop $(BIN_NAME).service && sudo mv ~/$(BIN_NAME) /usr/bin/$(BIN_NAME) && sudo systemctl start $(BIN_NAME).service"
