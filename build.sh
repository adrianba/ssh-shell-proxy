#!/bin/bash
set -euo pipefail

echo "Building ssh-shell-proxy for Windows x64..."
GOOS=windows GOARCH=amd64 go build -o ssh-shell-proxy-x64.exe .

echo "Building ssh-shell-proxy for Windows ARM64..."
GOOS=windows GOARCH=arm64 go build -o ssh-shell-proxy-arm64.exe .

echo "Done."
ls -lh ssh-shell-proxy-*.exe
