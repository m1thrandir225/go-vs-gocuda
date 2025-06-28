#include <cuda_runtime.h>
#include <device_launch_parameters.h>
#include <stdlib.h>
#include <stdio.h>
#include <cuda.h>
#include <math.h>
#include <cuda_runtime_api.h>

__global__ void matrix_multiplication(float *a, float *b, float *c, int width)
{
    int row = blockIdx.y * blockDim.y + threadIdx.y;
	int col = blockIdx.x * blockDim.x + threadIdx.x;

    if (row < width && col < width) {
        float sum = 0.0f;
        for (int k = 0; k < width; ++k) {
			sum += a[row * width + k] * b[k * width + col];
        }
        c[row * width + col] = sum;
    }
}

__global__ void tiled_matrix_mul(float *a, float *b, float *c, int width) {
    const int TILE_SIZE = 16;
   __shared__ float sA[TILE_SIZE][TILE_SIZE];
   __shared__ float sB[TILE_SIZE][TILE_SIZE];


   int bx = blockIdx.x;
   int by = blockIdx.y;
   int tx = threadIdx.x;
   int ty = threadIdx.y;

   int row = by * TILE_SIZE + ty;
   int col = bx * TILE_SIZE + tx;

   float c_value = 0.0f;

      for (int t = 0; t < (width + TILE_SIZE - 1) / TILE_SIZE; ++t) {
        if (row < width && (t * TILE_SIZE + tx) < width) {
            sA[ty][tx] = a[row * width + (t * TILE_SIZE + tx)];
        } else {
            sA[ty][tx] = 0.0f;
        }

        if ((t * TILE_SIZE + ty) < width && col < width) {
            sB[ty][tx] = b[(t * TILE_SIZE + ty) * width+ col];
        } else {
            sB[ty][tx] = 0.0f;
        }
        __syncthreads();

        for (int k = 0; k < TILE_SIZE; ++k) {
            c_value += sA[ty][k] * sB[k][tx];
        }

        __syncthreads();
    }

    if (row < width && col < width) {
        c[row * width + col] = c_value;
    }
}


extern "C" {
    __declspec(dllexport) void matrix_multiplication_wrapper(float *a, float *b, float *c, int size) {
		int total = size * size;

        float* d_a = NULL, * d_b = NULL, * d_c = NULL;
        int size_bytes = total * sizeof(float);

        // Allocate memory on the GPU
        cudaMalloc((void**)&d_a, size_bytes);
        cudaMalloc((void**)&d_b, size_bytes);
        cudaMalloc((void**)&d_c, size_bytes);
        
        // Copy data from host (CPU) to device (GPU)
        cudaMemcpy(d_a, a, size_bytes, cudaMemcpyHostToDevice);
        cudaMemcpy(d_b, b, size_bytes, cudaMemcpyHostToDevice);

        // Define grid and block dimensions
        dim3 threadsPerBlock(16, 16);
        dim3 numBlocks((size + threadsPerBlock.x - 1) / threadsPerBlock.x, (size + threadsPerBlock.y - 1) / threadsPerBlock.y);

        // Launch the kernel
        matrix_multiplication<<<numBlocks, threadsPerBlock>>>(d_a, d_b, d_c, size);

        // Copy the result back from device to host
        cudaMemcpy(c, d_c, size_bytes, cudaMemcpyDeviceToHost);

        // Free GPU memory
        cudaFree(d_a);
        cudaFree(d_b);
        cudaFree(d_c);
    }

    __declspec(dllexport) void tiled_matrix_multiplication_wrapper(float *a, float *b, float *c, int size) {
        int total = size * size;
        const int TILE_SIZE = 16;

        float *d_a = NULL, *d_b = NULL, *d_c = NULL;
        int size_bytes = total * sizeof(float);

        //Allocate GPU memory
        cudaMalloc((void**)&d_a, size_bytes);
        cudaMalloc((void**)&d_b, size_bytes);
        cudaMalloc((void**)&d_c, size_bytes);

        // Copy data from Host to device
        cudaMemcpy(d_a, a, size_bytes, cudaMemcpyHostToDevice);
        cudaMemcpy(d_b, b, size_bytes, cudaMemcpyHostToDevice);

        // Define grid and block dimensions
        int grid_dim_x = (size + TILE_SIZE - 1) / TILE_SIZE;
        int grid_dim_y = (size + TILE_SIZE - 1) / TILE_SIZE;
        dim3 gridDim(grid_dim_x, grid_dim_y);
        dim3 blockDim(TILE_SIZE, TILE_SIZE);

        //Launch kernel
        tiled_matrix_mul<<<gridDim, blockDim>>>(d_a, d_b, d_c, size);

        //Copy result back from device to host
        cudaMemcpy(c, d_c, size_bytes, cudaMemcpyDeviceToHost);

        //Free memory
        cudaFree(d_a);
        cudaFree(d_b);
        cudaFree(d_c);
    }
}

