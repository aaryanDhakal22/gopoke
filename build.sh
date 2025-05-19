#!/bin/bash

# Build script for Pokedex Go Application
# Usage: ./build.sh [command]
#   where command is one of:
#     build    - Build the application
#     test     - Run tests
#     run      - Build and run the application
#     clean    - Clean build artifacts
#     all      - Run clean, test, and build
#     help     - Show this help message

set -e  # Exit immediately if a command exits with a non-zero status

# Define variables
APP_NAME="gopoke"
BUILD_DIR="./build"
MAIN_PKG="./cmd/gopoke"
MODULE_NAME=$(grep "^module" go.mod | awk '{print $2}')
GO_VERSION=$(grep "^go" go.mod | awk '{print $2}')

# Color definitions
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Print a formatted message
print_message() {
    echo -e "${BLUE}[Build Script]${NC} $1"
}

# Print a success message
print_success() {
    echo -e "${GREEN}[Success]${NC} $1"
}

# Print an error message
print_error() {
    echo -e "${RED}[Error]${NC} $1"
}

# Print a warning message
print_warning() {
    echo -e "${YELLOW}[Warning]${NC} $1"
}

# Check Go installation
check_go() {
    if ! command -v go &> /dev/null; then
        print_error "Go is not installed or not in PATH"
        exit 1
    fi
    
    LOCAL_GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
    print_message "Using Go version: ${LOCAL_GO_VERSION}"
    print_message "Module requires Go version: ${GO_VERSION}"
    
    # Simple version check - you might want to implement a more robust version comparison
    if [[ "${LOCAL_GO_VERSION}" < "${GO_VERSION}" ]]; then
        print_warning "Your Go version might be older than required in go.mod"
    fi
}

# Clean build artifacts
clean() {
    print_message "Cleaning build artifacts..."
    
    if [ -d "$BUILD_DIR" ]; then
        rm -rf "$BUILD_DIR"
        print_success "Build directory cleaned"
    else
        print_message "Build directory does not exist, nothing to clean"
    fi
    
    # Clean test cache
    go clean -testcache
    print_success "Test cache cleaned"
}

# Run tests
run_tests() {
    print_message "Running tests..."
    
    # Run tests with verbose output and race detection
    go test -v -race ./...
    
    if [ $? -eq 0 ]; then
        print_success "All tests passed"
    else
        print_error "Tests failed"
        exit 1
    fi
}

# Build the application
build() {
    print_message "Building $APP_NAME..."
    
    # Create build directory if it doesn't exist
    mkdir -p "$BUILD_DIR"
    
    # Build with version information
    BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')
    GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
    
    go build -o "$BUILD_DIR/$APP_NAME" \
        -ldflags "-X '${MODULE_NAME}/internal/version.BuildTime=${BUILD_TIME}' -X '${MODULE_NAME}/internal/version.GitCommit=${GIT_COMMIT}'" \
        "$MAIN_PKG"
    
    if [ $? -eq 0 ]; then
        print_success "$APP_NAME built successfully"
    else
        print_error "Build failed"
        exit 1
    fi
}

# Run the application
run_app() {
    print_message "Running $APP_NAME..."
    
    # Check if the binary exists, if not, build it
    if [ ! -f "$BUILD_DIR/$APP_NAME" ]; then
        print_warning "Binary not found, building first..."
        build
    fi
    
    # Run the application
    "$BUILD_DIR/$APP_NAME"
}

# Show help message
show_help() {
    echo "Build script for $APP_NAME"
    echo
    echo "Usage: ./build.sh [command]"
    echo "Commands:"
    echo "  build    - Build the application"
    echo "  test     - Run tests"
    echo "  run      - Build and run the application"
    echo "  clean    - Clean build artifacts"
    echo "  all      - Run clean, test, and build"
    echo "  help     - Show this help message"
}

# Main script execution
check_go

case "$1" in
    build)
        build
        ;;
    test)
        run_tests
        ;;
    run)
        run_app
        ;;
    clean)
        clean
        ;;
    all)
        clean
        run_tests
        build
        print_success "All operations completed successfully"
        ;;
    help|"")
        show_help
        ;;
    *)
        print_error "Unknown command: $1"
        show_help
        exit 1
        ;;
esac

exit 0
