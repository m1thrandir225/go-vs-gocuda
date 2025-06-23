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

You can use the `Makefile` for this. (Windows, Mac & Linux).

## Performance Benchmarks

TODO:
