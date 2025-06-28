
# --- Variables ---

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

NVCC=nvcc
NVCCFLAGS=--shared -gencode=arch=compute_86,code=sm_86

DIST_DIR=.dist

CUDA_SRC_DIR=internals/cuda
NATIVE_CMD_DIR=cmd/native
CUDA_CMD_DIR=cmd/cuda

CUDA_SRC=$(CUDA_SRC_DIR)/matrix_mult.cu
CUDA_LIB_NAME=libmatrixmult.dll
CUDA_LIB_SRC_PATH=$(CUDA_SRC_DIR)/$(CUDA_LIB_NAME)

NATIVE_BIN=$(DIST_DIR)/native_benchmark.exe
CUDA_BIN=$(DIST_DIR)/cuda_benchmark.exe

# --- Build & Run Targets ---
.PHONY: all build run run-native run-cuda clean

all: build

build: $(NATIVE_BIN) $(CUDA_BIN)

$(DIST_DIR):
	@echo "Creating distribution directory..."
	mkdir -p $(DIST_DIR)

$(NATIVE_BIN): $(NATIVE_CMD_DIR)/*.go | $(DIST_DIR)
	@echo "Building Native Go benchmark..."
	$(GOBUILD) -o $@ ./$(NATIVE_CMD_DIR)

$(CUDA_BIN): $(CUDA_CMD_DIR)/*.go $(CUDA_LIB_SRC_PATH) | $(DIST_DIR)
	@echo "Building CUDA Go benchmark..."
	$(GOBUILD) -o $@ ./$(CUDA_CMD_DIR)
	@echo "Copying CUDA DLL to distribution directory..."
	cp $(CUDA_LIB_SRC_PATH) $(DIST_DIR)

$(CUDA_LIB_SRC_PATH): $(CUDA_SRC)
	@echo "Building CUDA shared library..."
	$(NVCC) -o $@ $(NVCCFLAGS) $<

run: run-native run-cuda

run-native: $(NATIVE_BIN)
	@echo "--- Running Native Benchmark ---"
	./$(NATIVE_BIN)
	@echo "--------------------------------"

run-cuda: $(CUDA_BIN)
	@echo "--- Running CUDA Benchmark ---"
	./$(CUDA_BIN)
	@echo "------------------------------"


# --- Cleanup ---
clean:
	@echo "Cleaning up build artifacts..."
	rm -rf $(DIST_DIR)
	rm -f $(CUDA_LIB_SRC_PATH)
	$(GOCLEAN)
