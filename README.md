# Go vs Go+Cuda

This project is a benchmark analysis of the possible speedup you could gain when utilizing GPU compute via CUDA.
The CUDA utilization is done via C bindings.

## How to Run

### Go+CUDA Benchmarks

Requirements: CUDA & GCC, Go 1.24+

The C binding's are done via a Dynamic-Link Library. Firstly you need to compile the CUDA code for your platform and then generate the needed libraries.

I've chosen to use a Dynamic-Link Library for simplicity of testing.

You can use the provided `Makefile` to compile for your system. (Windows & Linux only).

After that you can compile the Go binary for your system and run it alongside the built CUDA DLL.

### Go Benchmarks

Requirements: Go 1.24+

Run or compile the native package.

You can use the `Makefile` for this. (Windows).

## Performance Benchmarks

For my performance analysis I did test it on 3 machines.

1. RTX 3070 + Ryzen 7 3700x
2. Macbook M2 MAX (CPU Only Comparison)

### RTX 3070 + Ryzen 7 3700x

| Size      | CPU-Basic      | CPU-Parallel | CPU-Parallel-Worker-Pool | GPU-Basic              | GPU-Tiled             |
| --------- | -------------- | ------------ | ------------------------ | ---------------------- | --------------------- |
| 512x512   | 457.8794ms     | 39.4346ms    | 38.1983ms                | 75.9094ms(73.1074ms)   | 4.0501ms(3.0408ms)    |
| 1024x1024 | 3.7636204s     | 360.9949ms   | 358.5513ms               | 83.2571ms(79.5978ms)   | 15.553ms(11.6207ms)   |
| 2048x2048 | 40.3201839s    | 4.6126668s   | 4.6643551s               | 152.9811ms(140.8297ms) | 85.9288ms(74.6713ms)  |
| 4096x4096 | 14m58.8036904s | 2m6.8188304s | 2m31.7133689s            | 585.6444ms(540.149ms)  | 494.473ms(448.6114ms) |

### Macbook M2 MAX

| Size      | CPU-Basic      | CPU-Parallel   | CPU-Parallel-Worker-Pool | GPU-Basic | GPU-Tiled |
| --------- | -------------- | -------------- | ------------------------ | --------- | --------- |
| 512x512   | 489.019083ms   | 22.304417ms    | 22.450125ms              | /         | /         |
| 1024x1024 | 4.1104115      | 213.212875ms   | 217.149334ms             | /         | /         |
| 2048x2048 | 57.987378833   | 8.697457584    | 8.903120709              | /         | /         |
| 4096x4096 | 9m21.453859875 | 1m12.918281041 | 1m13.366087833           | /         | /         |
