.PHONY: build clean

build:
	@echo "Building Go binary with optimization flags..."
	go build -ldflags="-s -w" -o box.exe cmd/app/main.go
	@echo "Build complete! Final size:"
	@du -h box.exe
	@echo "Compressing executable with UPX..."
	upx --best --lzma box.exe
	@echo "Build complete! Final size:"
	@du -h box.exe

clean:
	@echo "Cleaning build artifacts..."
	rm -f box.exe
	@echo "Clean complete"