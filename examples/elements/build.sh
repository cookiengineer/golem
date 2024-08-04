#!/bin/bash

GOROOT=$(go env GOROOT);

env GOOS=js GOARCH=wasm go build -o "public/main.wasm" main.go;
cp "${GOROOT}/misc/wasm/wasm_exec.js" "public/wasm_exec.js";

