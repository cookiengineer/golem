#!/bin/bash

GOROOT=$(go env GOROOT);
ROOT=$(pwd);

env GOOS=js GOARCH=wasm go build -o "${ROOT}/public/main.wasm" main.go;
cp "${GOROOT}/misc/wasm/wasm_exec.js" "${ROOT}/public/wasm_exec.js";

go run "${ROOT}/serve.go";

