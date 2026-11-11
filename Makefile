# ==========================================
# L4D2 Mod Manager (l4d2mm) Makefile
# ==========================================

APP_NAME := l4d2mm
BUILD_DIR := build
SRC_ENTRY := .

.PHONY: all clean \
        windows linux \
        amd64 arm64 \
        windows-amd64 windows-arm64 \
        linux-amd64 linux-arm64

# Build all platforms and architectures
all: windows linux

# ==========================================
# 1. Build by OS (all architectures for the OS)
# ==========================================
windows: windows-amd64 windows-arm64
linux: linux-amd64 linux-arm64

# ==========================================
# 2. Build by CPU architecture (cross-platform)
# ==========================================
amd64: windows-amd64 linux-amd64
arm64: windows-arm64 linux-arm64

# ==========================================
# 3. Individual build targets (single OS + single arch)
# ==========================================
windows-amd64:
	@echo "Building for Windows (amd64)..."
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)_windows_amd64.exe $(SRC_ENTRY)

windows-arm64:
	@echo "Building for Windows (arm64)..."
	GOOS=windows GOARCH=arm64 go build -o $(BUILD_DIR)/$(APP_NAME)_windows_arm64.exe $(SRC_ENTRY)

linux-amd64:
	@echo "Building for Linux (amd64)..."
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)_linux_amd64 $(SRC_ENTRY)

linux-arm64:
	@echo "Building for Linux (arm64)..."
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(APP_NAME)_linux_arm64 $(SRC_ENTRY)

# ==========================================
# Clean build artifacts
# ==========================================
clean:
	@echo "Cleaning build directory..."
	rm -rf $(BUILD_DIR)