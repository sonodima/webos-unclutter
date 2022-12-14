BINARY_NAME = wosu
BIN_DIR = bin


###############################################
## Do not edit below this line
###############################################

BINARY_TARGET					= $(BIN_DIR)/$(BINARY_NAME)
MACHO_ARM64_TARGET		= $(BINARY_TARGET)-arm64
MACHO_X86_64_TARGET		= $(BINARY_TARGET)-x86_64

BUILD_FLAGS = -ldflags "-s -w"

# Enable experimental features for multi-arch container images
export DOCKER_CLI_EXPERIMENTAL=enabled


# Get the OS of the machine that is building the program
ifeq ($(OS),Windows_NT)
	BLDOS = windows
else
	BLDOS = $(shell uname -s | tr '[:upper:]' '[:lower:]')
endif

# If the GOOS environment variable is not set, default to the current OS
ifeq ($(GOOS),)
GOOS = $(BLDOS)
endif


.SILENT: build
build:
ifeq ($(GOOS), windows)

  ################
  ## Windows
  ################

	echo "πͺ Building for Windows"
	go build -o $(BINARY_TARGET).exe $(BUILD_FLAGS) .

else ifeq ($(GOOS), darwin)

  ################
  ## macOS
  ################

ifeq ($(BLDOS), darwin)
  # If the current OS is macOS, we can build a U2B binary using LIPO.
	echo "π Building U2B for macOS"

	echo "π© Compiling for arm64"
	GOARCH=arm64 go build $(BUILD_FLAGS) -o $(MACHO_ARM64_TARGET) .
	echo "π© Compiling for x86_64"
	GOARCH=amd64 go build $(BUILD_FLAGS) -o $(MACHO_X86_64_TARGET) .

  # Create a FAT binary (contains both arm64 and x86_64)
	echo "π¦ Merging build outputs into a FAT binary"
	lipo -create -output $(BINARY_TARGET) $(MACHO_ARM64_TARGET) $(MACHO_X86_64_TARGET)

  #Β Remove the THIN binaries
	echo "π§Ό Removing standalone arm64 and x86_64 binaries"
	rm $(MACHO_ARM64_TARGET) $(MACHO_X86_64_TARGET)
else
  # If the current OS is not macOS, we must build a single THIN binary for the target architecture.
	echo "π Building for macOS"
	go build $(BUILD_FLAGS) -o $(BINARY_TARGET) .
endif

else

  ################
  ## Linux
  ################

	echo "π§ Building for Linux"
	go build -o $(BINARY_TARGET) $(BUILD_FLAGS) .

endif
	
	echo "β \033[1;32mCompilation succeeded\033[0m"
	

.SILENT: run
.PHONY: run
run:
	go run .


.SILENT: docker
.PHONY: docker
docker:
	echo "π§ Building Docker image"
	docker buildx create --use --name=wosu --node=wosu
	docker buildx build --output "type=docker,push=false" --tag wosu:latest .
	docker run -it -d -p 53:53/udp wosu:latest
  # docker buildx build --platform linux/amd64,linux/arm64 --push --tag wosu:latest .


.SILENT: clean
.PHONY: clean
clean:
	echo "π§Ό Removing object files from package source directories"
	go clean
	echo "π§Ό Removing build outputs"
	rm -f $(BIN_DIR)/*
	echo "β \033[1;32mClean succeeded\033[0m"


.SILENT: help
.PHONY: help
help:
	echo
	echo "\033[1;32mUSAGE:\033[0m"
	echo "  [vars] make [target]"
	echo
	echo "\033[1;32mTARGETS:\033[0m"
	echo "  build: Build the program"
	echo "  run: Run the program"
	echo "  docker: Build and run the docker image"
	echo "  clean: Clean the build directory"
	echo "  help: Show this help message"
	echo
	echo "\033[1;32mVARS:\033[0m"
	echo "  GOOS: The OS to build for (default: $(GOOS))"
