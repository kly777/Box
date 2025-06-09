.PHONY: build clean version cross-build lint fmt tidy help

# 获取当前git版本号
VERSION := $(shell git describe --tags 2>/dev/null || echo "v0.0.0")

run:
	@echo "Running..."
	go run cmd/app/main.go



# 默认构建目标
build:
	@echo "Building Go binary (version: ${VERSION}) with optimization flags..."
	go build -ldflags="-s -w -H windowsgui" -o Box.exe cmd/app/main.go
	# @echo "Compressing executable with UPX..."
	# upx -6 --fast --lzma Box.exe
	# @echo "Build complete! Final size:"
	@du -h Box.exe

buildDBG:
	@echo "Building Go binary (version: ${VERSION}) with debug flags..."
	go build -ldflags="-s -w -X main.version=${VERSION}" -o BoxDBG.exe cmd/app/main.go
	@du -h BoxDBG.exe

# 清理构建产物
clean:
	@echo "Cleaning build artifacts..."
	rm -f Box.exe Box-linux
	@echo "Clean complete"

# 显示版本信息
version:
	@echo "Current version: ${VERSION}"

# 交叉编译Linux版本
cross-build:
	@echo "Cross-compiling for Linux (version: ${VERSION})..."
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X main.version=${VERSION}" -o Box-linux cmd/app/main.go
	@echo "Linux build complete! File: Box-linux"

# 代码质量检查
lint:
	@if ! command -v golangci-lint >/dev/null; then \
		echo "golangci-lint not found. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	@golangci-lint run ./...

# 代码格式化检查
fmt:
	@echo "Checking code formatting..."
	@if [ -n "$$(go fmt ./...)" ]; then \
		echo "Code formatting issues found. Run 'go fmt ./...' to fix"; \
		exit 1; \
	else \
		echo "Code is properly formatted"; \
	fi

# 依赖管理
tidy:
	@echo "Tidying module dependencies..."
	go mod tidy -v

