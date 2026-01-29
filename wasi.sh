#!/bin/sh

outname=./xlsx-ls-sheets.wasm
mainpat=./cmd/xlsx-ls-sheets/main.go

build_std() {
	GOOS=wasip1 GOARCH=wasm go \
		build \
		-o "${outname}" \
		-ldflags="-s -w" \
		"${mainpat}"
}

build_tiny() {
	tinygo \
		build \
		-o "${outname}" \
		-target=wasip1 \
		-opt=z \
		-no-debug \
		"${mainpat}"
}

build_std
#build_tiny
