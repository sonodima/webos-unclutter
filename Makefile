BINARY_NAME = wosu
BIN_DIR = bin


###############################################
## Do not edit below this line
###############################################

BINARY_TARGET					= $(BIN_DIR)/$(BINARY_NAME)
MACHO_ARM64_TARGET		= $(BINARY_TARGET)-arm64
MACHO_X86_64_TARGET		= $(BINARY_TARGET)-x86_64

BUILD_FLAGS = -ldflags "-s -w"
UNAME = $(shell uname -s)

export DOCKER_CLI_EXPERIMENTAL=enabled


.SILENT: build
build:
	echo "🔧 Build started"
	
  # macOS universal binaries contain both x86_64 and arm64 binaries
ifeq ($(UNAME),Darwin)
	echo "🍎 Building U2B for macOS"

	echo "🔩 Compiling for arm64"
	GOARCH=arm64 go build $(BUILD_FLAGS) -o $(MACHO_ARM64_TARGET) .
	echo "🔩 Compiling for x86_64"
	GOARCH=amd64 go build $(BUILD_FLAGS) -o $(MACHO_X86_64_TARGET) .

	echo "📦 Merging build outputs into a FAT binary"
	lipo -create -output $(BINARY_TARGET) $(MACHO_ARM64_TARGET) $(MACHO_X86_64_TARGET)

	echo "🧼 Removing standalone arm64 and x86_64 binaries"
	rm $(MACHO_ARM64_TARGET) $(MACHO_X86_64_TARGET)
else
	go build $(BUILD_FLAGS) -o $(BINARY_TARGET) .	
endif
	
	echo "✅ \033[1;32mCompilation succeeded\033[0m"
	

.SILENT: run
.PHONY: run
run:
	go run .


.SILENT: docker
.PHONY: docker
docker:
	echo "🔧 Building Docker image"
	docker buildx create --use --name=wosu --node=wosu
	docker buildx build --output "type=docker,push=false" --tag wosu:latest .
	docker run -it -d -p 53:53/udp wosu:latest
  # docker buildx build --platform linux/amd64,linux/arm64 --push --tag wosu:latest .


.SILENT: clean
.PHONY: clean
clean:
	echo "🧼 Removing object files from package source directories"
	go clean
	echo "🧼 Removing build outputs"
	rm -f $(BIN_DIR)/*
	echo "✅ \033[1;32mClean succeeded\033[0m"


.SILENT: help
.PHONY: help
help:
	echo
	echo "\033[1;32mUSAGE:\033[0m"
	echo "  make [target]"
	echo
	echo "\033[1;32mTARGETS:\033[0m"
	echo "  build: Build the program"
	echo "  run: Run the program"
	echo "  docker: Build and run the docker image"
	echo "  clean: Clean the build directory"
	echo "  help: Show this help message"
