SHELL=cmd


# Application version from git tag (latest tag)
APP_VERSION=$(shell git describe --tags --abbrev=0)

# Git commit hash
GIT_COMMIT=$(shell git rev-parse --short HEAD)

# Application name
APP_NAME=tarkov-server-checker

BIN_NAME=$(APP_NAME).exe

BIN_PATH=.\bin

# Application entrypoint
ENTRYPOINT=.\main.go

# Build path
BUILD_PATH=$(BIN_PATH)\$(BIN_NAME)

# LD flags for GUI applications on Windows (no console) and version info
LDFLAGS=-ldflags="-s -w -H windowsgui -X cmd.Version=$(APP_VERSION) -X cmd.Commit=$(GIT_COMMIT)"

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build $(LDFLAGS)
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

all:
	echo "Please use 'make build' or 'make run' to build or run the application."

# Build for Windows
build:
	echo "Building for Windows amd64 (64bit) ..."
	$(GOBUILD) -o $(BUILD_PATH) -v $(ENTRYPOINT)

# Clean the build
clean:
	echo "Cleaning the build..."
	$(GOCLEAN) -i $(ENTRYPOINT)
	del $(BUILD_PATH)

# Clean and build for Windows
cleanbuild: clean build

# Run the application
run:
	echo "Running the application..."
	go run $(ENTRYPOINT)