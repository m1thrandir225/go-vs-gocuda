#include <cuda_runtime.h>
#include <device_launch_parameters.h>
#include <stdlib.h>
#include <stdio.h>
#include <cuda.h>
#include <math.h>
#include <cuda_runtime_api.h>
#define MAX_TILE_DIM 32

__global__ void matrix_multiplication_tiled(float* a, float* b, float* c, int width, int tileDim) {
    __shared__ float sA[MAX_TILE_DIM][MAX_TILE_DIM];
    __shared__ float sB[MAX_TILE_DIM][MAX_TILE_DIM];

    int tx = threadIdx.x;
    int ty = threadIdx.y;
    int row = blockIdx.y * tileDim + ty;
    int col = blockIdx.x * tileDim + tx;


    float sum = 0.0f;
    int numTiles = (width + tileDim - 1) / tileDim;

    for (int tile = 0; tile < numTiles; ++tile) {
        int a_load_row = row;
        int a_load_col = tile * tileDim + tx;
        if ((a_load_row < width & a_load_col < width & ty < tileDim & tx) < tileDim) {
            sA[ty][tx] = a[a_load_row * width + a_load_col];
        }
        else {
            sA[ty][tx] = 0.0f;
        }
        int b_load_row = tile * tileDim + ty;
        int b_load_col = col;
        if ((b_load_row < width & b_load_col < width & ty < tileDim & tx) < tileDim) {
            sB[ty][tx] = b[b_load_row * width + b_load_col];
        }
        else {
            sB[ty][tx] = 0.0f;
        }
        __syncthreads();

        if (ty < tileDim && tx < tileDim) {
            for (int k = 0; k < tileDim; ++k) {
                if ((tile * tileDim + k) < width) {
                    sum += sA[ty][k] * sB[k][tx];
                }
            }
        }
        __syncthreads();
    }
    if (row < width && col < width) {
        c[row * width + col] = sum;
    }
}

extern "C" {
    __declspec(dllexport) void matrix_multiplication_tiled_wrapper(float *a, float *b, float *c, int size)
    {
        int total = size * size;
        float *d_a = NULL, *d_b = NULL, *d_c = NULL;
        int size_bytes = total * sizeof(float);

        cudaMalloc((void**)&d_a, size_bytes);
	    cudaMalloc((void**)&d_b, size_bytes);
	    cudaMalloc((void**)&d_c, size_bytes);


	    cudaMemcpy(d_a, a, size_bytes, cudaMemcpyHostToDevice);
	    cudaMemcpy(d_b, b, size_bytes, cudaMemcpyHostToDevice);

        cudaEvent_t start, stop;
        cudaEventCreate(&start);
        cudaEventCreate(&stop);

        int tile_dims[] = { 8, 12, 16, 20, 24, 26, 28, 30, 32};
        int num_tile = sizeof(tile_dims) / sizeof(tile_dims[0]);

        float best_tiled_dim_ms = -1.0f;
        int best_tile_dim = 0;

        for (int i = 0; i < num_tile; i++) {
            int current_tile = tile_dims[i];

            if(current_tile > MAX_TILE_DIM) {
                printf("Exceeded max tile dim limit");
                continue;
            }

            printf("Testing dimension: %d\n", current_tile);
            dim3 threadsPerBlock(current_tile, current_tile);

            dim3 numBlocksTiled(( + current_tile - 1) / current_tile, (size + current_tile - 1) / current_tile);
            cudaEventRecord(start);
		    matrix_multiplication_tiled<<<numBlocksTiled, threadsPerBlock>>>(d_a, d_b, d_c, size, current_tile);

            
            cudaGetLastError();
            cudaEventRecord(stop);
            cudaEventSynchronize(stop);
		    float current_tiled_time_ms;

		    cudaEventElapsedTime(&current_tiled_time_ms, start, stop);

            if(best_tiled_dim_ms < 0 || current_tiled_time_ms < best_tiled_dim_ms) {
                best_tiled_dim_ms = current_tiled_time_ms;
                best_tile_dim = current_tile;
                cudaMemcpy(c, d_c, size_bytes, cudaMemcpyDeviceToHost);
		    }
		        cudaMemset(d_c, 0, size_bytes);
        }
        printf("Best tile dimension: %d with time: %f ms\n", best_tile_dim, best_tiled_dim_ms);

        cudaEventDestroy(start);
        cudaEventDestroy(stop);
	    cudaFree(d_a);
        cudaFree(d_b);
        cudaFree(d_c);
    }
}