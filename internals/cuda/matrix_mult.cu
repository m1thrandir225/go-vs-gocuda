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


extern "C" {
    __declspec(dllexport) void matrix_multiplication_wrapper(float* a, float* b, float* c, int size) {
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
        matrix_multiplication<<<numBlocks, threadsPerBlock >>>(d_a, d_b, d_c, size);

        // Copy the result back from device to host
        cudaMemcpy(c, d_c, size_bytes, cudaMemcpyDeviceToHost);

        // Free GPU memory
        cudaFree(d_a);
        cudaFree(d_b);
        cudaFree(d_c);
    }
}

