#!/bin/sh

outname=./xlsx-ls-sheets.tiny.wasm
binname=./xlsx-ls-sheets.tiny.cwasm
mainpat=./cmd/xlsx-ls-sheets/main.go

build_tiny() {
	echo 'building the wasi byte code(portable)...'
	echo 'this may took minutes(e.g., 60s on M3 Max with 48GB ram)'
	\time -l tinygo \
		build \
		-o "${outname}" \
		-target=wasip1 \
		-opt=z \
		-no-debug \
		"${mainpat}"
}

aot_wasmtime() {
	echo 'compiling the wasi byte code to native code(not portable; platform/runtime version dependent)...'
	echo 'this may took minutes(e.g., 120s on M3 Max with 48GB ram)'
	\time -l wasmtime \
		compile \
		"${outname}" \
		-o "${binname}"
}

run_aot_wasmtime() {
	cat ./input.xlsx |
		\time -l wasmtime \
			run \
			--allow-precompiled \
			-W max-memory-size=16777216 \
			"${binname}"
}

test -f "${outname}" || build_tiny
test -f "${binname}" || aot_wasmtime

run_aot_wasmtime
