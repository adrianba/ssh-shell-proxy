#!/bin/bash
set -euo pipefail

VERSION="${1:-dev}"

echo "Building ssh-shell-proxy v${VERSION} for Windows x64..."
GOOS=windows GOARCH=amd64 go build -ldflags="-X main.version=${VERSION}" -o ssh-shell-proxy-x64.exe .

echo "Building ssh-shell-proxy v${VERSION} for Windows ARM64..."
GOOS=windows GOARCH=arm64 go build -ldflags="-X main.version=${VERSION}" -o ssh-shell-proxy-arm64.exe .

echo "Done."
ls -lh ssh-shell-proxy-*.exe
