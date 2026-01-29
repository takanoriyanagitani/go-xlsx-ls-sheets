# go-xlsx-ls-sheets
Prints the sheet names of the xlsx file using wasi

## Simple Benchmark

### Env

#### Version info

- wasmtime: 41.0.1(2026-01-26)
- wasmer: 5.0.6(unable to upgrade)
- wasmedge: 0.16.1
- wazero: v1.11.0
- wasmi_cli: 1.0.7

#### Host

- OS: macOS 26
- Processor: M3 Max
- RAM: 48 GB

### Binary size(aot: for arm64-unknown-darwin)

| type         | go    | wasm-opt | size  | ratio  |          |
|:------------:|:-----:|:--------:|:-----:|:------:|:---------|
| wasmtime aot | std   | yes      | 19. M |  792%  | ********
| wasmtime aot | std   | no       | 20. M |  833%  | ********
| wasmtime aot | tiny  | yes      | 8.2 M |  342%  | ***
| wasmtime aot | tiny  | no       | 8.3 M |  346%  | ***
| wasmedge aot | std   | yes      | 11. M |  458%  | *****
| wasmedge aot | std   | no       | 12. M |  500%  | *****
| wasmedge aot | tiny  | yes      | 5.2 M |  217%  | **
| wasmedge aot | tiny  | no       | 5.3 M |  221%  | **
| wasm         | std   | yes      | 6.3 M |  263%  | ***
| wasm         | std   | no       | 6.9 M |  288%  | ***
| wasm         | tiny  | yes      | 2.4 M | (100%) | *
| wasm         | tiny  | no       | 2.5 M |  104%  | *

### Tiny Go

#### Run

| runtime           | wasm-opt | real | user | sys  | rss   | remarks |
|:-----------------:|:--------:|:----:|:----:|:----:|:-----:|:--------|
| wasmtime aot      | yes      | 0.01 | 0.00 | 0.00 |  22 M |
| wasmtime aot      | no       | 0.01 | 0.00 | 0.00 |  22 M |
| wasmedge aot      | yes      | 0.04 | 0.02 | 0.01 |  47 M |
| wasmedge aot      | no       | 0.02 | 0.02 | 0.01 |  47 M |
| wasmtime w/ cache | yes      | 0.03 | 0.02 | 0.00 |  44 M |
| wasmtime w/ cache | no       | 0.02 | 0.03 | 0.00 |  45 M |
| wasmtime no cache | yes      | 122. | 121. | 2.16 |  16 G | Compile time mem usage is high
| wasmtime no cache | no       | 121. | 119. | 2.27 |  16 G | Compile time mem usage is high
| wazero            | yes      | 10.1 | 10.2 | 0.07 | 564 M |
| wazero            | no       | 10.2 | 10.3 | 0.07 | 560 M |
| wasmer            | yes      | 0.97 | 0.89 | 0.02 |  41 M |
| wasmer            | no       | 0.91 | 0.89 | 0.01 |  42 M |
| wasmedge(1st)     | no       | 0.51 | 0.37 | 0.04 |  99 M |
| wasmedge(2nd)     | no       | 0.40 | 0.38 | 0.01 |  99 M |
| wasmedge(1st)     | yes      | 0.41 | 0.39 | 0.01 |  99 M |
| wasmedge(2nd)     | yes      | 0.38 | 0.37 | 0.00 |  99 M |
| wasmi_cli         | yes      | 0.11 | 0.10 | 0.00 |  21 M |
| wasmi_cli         | no       | 0.15 | 0.11 | 0.01 |  23 M |

#### Compile(AOT)

| runtime           | wasm-opt | real | user | sys  | rss   |
|:-----------------:|:--------:|:----:|:----:|:----:|:-----:|
| wasmedge(default) | yes      | 290. | 289. | 1.16 | 5.2 G |
| wasmedge(default) | no       | 292. | 291. | 1.24 | 5.2 G |
| wasmedge(opt 1)   | yes      | 263. | 262. | 1.35 | 7.4 G |
| wasmedge(opt 1)   | no       | 262. | 261. | 1.36 | 7.3 G |
| wasmtime          | yes      | 122. | 120. | 2.38 | 16. G |
| wasmtime          | no       | 121. | 120. | 1.97 | 16. G |

---

### Standard Go

#### Run

| runtime           | wasm-opt | real | user | sys  | rss   | rss ratio |
|:-----------------:|:--------:|:----:|:----:|:----:|:-----:|:---------:|
| wasmtime aot      | yes      | 0.02 | 0.02 | 0.00 |  32 M |  (100%)   |
| wasmtime aot      | no       | 0.02 | 0.01 | 0.00 |  33 M |   103%    |
| wasmedge aot      | yes      | 0.04 | 0.03 | 0.00 |  91 M |   284%    |
| wasmedge aot      | no       | 0.04 | 0.05 | 0.02 |  93 M |   291%    |
| wasmtime w/ cache | yes      | 0.05 | 0.04 | 0.00 |  80 M |   250%    |
| wasmtime w/ cache | no       | 0.05 | 0.05 | 0.00 |  83 M |   259%    |
| wasmtime no cache | yes      | 0.27 | 2.76 | 0.07 | 339 M |  1059%    |
| wasmtime no cache | no       | 0.28 | 2.92 | 0.07 | 323 M |  1009%    |
| wazero            | yes      | 1.14 | 1.16 | 0.04 | 285 M |   891%    |
| wazero            | no       | 1.21 | 1.25 | 0.03 | 311 M |   972%    |
| wasmer            | yes      | 0.20 | 0.18 | 0.01 |  88 M |   275%    |
| wasmer            | no       | 0.24 | 0.20 | 0.02 |  96 M |   300%    |
| wasmedge          | yes      | 0.54 | 0.48 | 0.03 | 188 M |   588%    |
| wasmedge          | no       | 0.54 | 0.50 | 0.02 | 217 M |   678%    |
| wasmi_cli         | yes      | 0.16 | 0.11 | 0.02 |  40 M |   125%    |
| wasmi_cli         | no       | 0.11 | 0.10 | 0.00 |  41 M |   128%    |

#### Compile(AOT)

| runtime           | wasm-opt | real | user | sys  | rss   |
|:-----------------:|:--------:|:----:|:----:|:----:|:-----:|
| wasmedge(default) | yes      | 79.5 | 79.0 | 0.49 | 1.9 G |
| wasmedge(default) | no       | 82.8 | 81.2 | 0.55 | 2.0 G |
| wasmedge(opt 1)   | yes      | 62.3 | 61.8 | 0.48 | 2.0 G |
| wasmedge(opt 1)   | no       | 58.7 | 58.3 | 0.44 | 1.9 G |
| wasmtime          | yes      | 0.36 | 2.94 | 0.07 | 309 M |
| wasmtime          | no       | 0.33 | 2.71 | 0.07 | 317 M |

