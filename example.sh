#!/bin/sh

run_wasi() {
	cat ./input.xlsx |
		wasmi_cli \
			./xlsx-ls-sheets.wasm
}

run_wasi
