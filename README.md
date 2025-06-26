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

For my performance analysis I did test it on 3 machines.

1. RTX 3070 + Ryzen 7 3700x
2. Macbook M2 MAX (CPU Only Comparison)
3. ...

The CPU Mode has 3 different benchmarks on it i.e making it faster.

The simplest one is just a for loop of matrix multiplication.

The `Parallel` mode is spawing Go-routes and using a `sync.WaitGroup` simillar to how we would do it in CUDA.

The third mode called - `Worker Pool` makes jobs as much CPU threads we have available on our system, so we can maximize performance and don't have any performance loss.

### RTX 3070 vs Ryzen 7 3700x

**CPU**

```bash
Native Basic - 493.4545ms (Verification: True)
Native Parallel Mode - 38.6415ms  (Verification: True)
Native Parallel Worker Pool - 37.6499ms (Verification: True)
```

In most of my tests the difference between Parallel Mode and Parallel Worker Mode ~ 2% difference in benefit of the most optimized version i.e the Paralell Worker Mode. But the difference between the Parallel modes and the Basic mode is about a 92% performance gain.

**GPU**
