# This Makefile is used to build and run the tarkov-server-checker application. 

# Set the shell to PowerShell
SHELL=powershell

# Define the paths to the application's resources
ICON_PATH=.\resources\icon\icon.ico
BIN_PATH=.\bin
TOOLS_PATH=.\tools
ENTRYPOINT=.\cmd

# Define the version information
VERSIONINFO_PATH=$(ENTRYPOINT)\tarkov-server-checker\versioninfo.json
VERSIONINFO_MANAGER_PATH=$(TOOLS_PATH)\versioninfo_manager
VERSIONINFO_MANAGER=$(GOCMD) run $(VERSIONINFO_MANAGER_PATH)
VERSIONINFO_GET_VERSION=$(VERSIONINFO_MANAGER) get -j $(VERSIONINFO_PATH)

# Define the application version (major.minor.patch.build)
# The build is auto-incremented by the installer script (installer.nsi). From git tags.
APP_VERSION=$(shell git describe --tags --abbrev=0)

# Define the application name and related paths
APP_NAME=tarkov-server-checker
BIN_NAME=$(APP_NAME).exe
BUILD_PATH=$(BIN_PATH)\$(APP_VERSION)\$(BIN_NAME)

# Define the installer name and related paths
INSTALLER_NAME=$(APP_NAME)-$(APP_VERSION)-setup.exe
INSTALLER_BUILD_PATH=$(BIN_PATH)\$(APP_VERSION)\$(INSTALLER_NAME)

# Define the optimization flags for the compiler
LDFLAGS=-ldflags="-s -w -H=windowsgui"

# Define the Go parameters
GOCMD=go
GOGENERATE=$(GOCMD) generate
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
NSIS=makensis.exe

# Define the default target
all:
	@echo "Please use 'make build' or 'make run' to build or run the application."

# Define the build target
build: build-app build-installer

# Define the target to create the build directory
build-mkdir:
    $(shell if not exist $(BIN_PATH)\$(APP_VERSION) mkdir $(BIN_PATH)\$(APP_VERSION))

# Define the target to build the application binary
build-app: clean build-mkdir clean-app
	$(GOGENERATE) $(ENTRYPOINT)\tarkov-server-checker
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_PATH) -v $(ENTRYPOINT)\tarkov-server-checker

# Define the target to build the installer
build-installer: build-mkdir get-version clean-installer build-nsis

# Define the target to build the NSIS installer
build-nsis: build-mkdir
	$(NSIS) /DVERSION=${VERSION} /DBUILD_PATH=$(BUILD_PATH) /DBIN_NAME=$(BIN_NAME) /DICON_PATH=$(ICON_PATH) /DINSTALLER_BUILD_PATH=$(INSTALLER_BUILD_PATH) .\installer.nsi

# Define the target to clean the build and syso files, versioninfo files
clean: clean-app clean-installer clean-syso

# Define the target to clean the application build
clean-app: clean-syso
    $(shell if exist $(BUILD_PATH) del $(BUILD_PATH))

# Define the target to clean the installer
clean-installer:
    $(shell if exist $(INSTALLER_BUILD_PATH) del $(INSTALLER_BUILD_PATH))

# Define the target to clean syso files for the application (version info)
clean-syso:
    $(shell if not exist $(SYSO_FILE_PATH) del $(SYSO_FILE_PATH))

# Define the target to run the application
run:
	$(GOCMD) run $(ENTRYPOINT)\tarkov-server-checker

# Define the target to get the version
get-version:
	$(eval VERSION = $(shell $(VERSIONINFO_GET_VERSION)))