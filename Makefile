# Makefile for building and running the tarkov-server-checker application

SHELL=cmd

# Paths
ICON_PATH=.\resources\icon\icon.ico
BIN_PATH=.\bin
ENTRYPOINT=.\cmd
VERSIONINFO_PATH=$(ENTRYPOINT)\tarkov-server-checker\versioninfo.json

# Application version (major.minor.patch.build) - build is auto-incremented by the installer script (installer.nsi). From git tags.
APP_VERSION=$(shell git describe --tags --abbrev=0)

# Application name
APP_NAME=tarkov-server-checker
BIN_NAME=$(APP_NAME).exe
BUILD_PATH=$(BIN_PATH)\$(APP_VERSION)\$(BIN_NAME)
ICON_PATH=.\resources\icon\icon.ico
# Installer name
INSTALLER_NAME=$(APP_NAME)-$(APP_VERSION)-setup.exe
INSTALLER_BUILD_PATH=$(BIN_PATH)\$(APP_VERSION)\$(INSTALLER_NAME)

# Optimization flags for the compiler, remove debug information and disable assertions (for smaller binaries). Also remove the console window.
LDFLAGS=-ldflags="-s -w -H=windowsgui"

# Go parameters
GOCMD=go
GOGENERATE=$(GOCMD) generate
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
NSIS=makensis.exe
VERSIONINFO_MANAGER=$(GOCMD) run $(ENTRYPOINT)\versioninfo_manager

all:
	@echo "Please use 'make build' or 'make run' to build or run the application."

build: run-versioninfo_manager build-tarkov-server-checker

# Build for Windows
build-tarkov-server-checker:
	@echo "Building for Windows amd64 (64bit) ..."
	mkdir $(BIN_PATH)\$(APP_VERSION)
	$(GOGENERATE) $(ENTRYPOINT)\tarkov-server-checker
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_PATH) -v $(ENTRYPOINT)\tarkov-server-checker
	$(NSIS) /DVERSION=1.0.0.0 /DBUILD_PATH=$(BUILD_PATH) /DBIN_NAME=$(BIN_NAME) /DICON_PATH=$(ICON_PATH) /DINSTALLER_BUILD_PATH=$(INSTALLER_BUILD_PATH) .\installer.nsi

# Clean the build and syso files, versioninfo files
clean: clean-build clean-syso

# Clean the build
clean-build:
	@echo "Cleaning the build..."
	rmdir /s /q $(BIN_PATH)

# Clean syso files for the application (version info)
clean-syso:
	@echo "Cleaning syso files..."
	del $(ENTRYPOINT)\tarkov-server-checker\*.syso

# Run the application
run:
	@echo "Running the application..."
	$(GOCMD) run $(ENTRYPOINT)\tarkov-server-checker

run-versioninfo_manager:
	@echo "Running versioninfo manager..."
	$(VERSIONINFO_MANAGER) --gitTag $(APP_VERSION) -j $(VERSIONINFO_PATH)
