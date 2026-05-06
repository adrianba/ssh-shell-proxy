# ssh-shell-proxy

A small Windows executable that launches a WSL Debian shell via `wsl.exe`.

## Usage

**Interactive shell:**

```
ssh-shell-proxy.exe
```

Opens a WSL Debian shell in your home directory.

**Run a command:**

```
ssh-shell-proxy.exe -c "ls -la"
```

Runs the command inside WSL Debian and exits.

**Any other arguments** will print an error to stderr.

## Prerequisites

### Install Go

Install or update Go using [webi](https://webinstall.dev/golang/) (no sudo required):

```bash
curl -sS https://webi.sh/golang | sh
```

Then activate it in your current terminal:

```bash
source ~/.config/envman/PATH.env
```

To update Go later:

```bash
webi golang@stable
```

## Building

Run the build script to compile for both Windows x64 and ARM64:

```bash
./build.sh
```

This produces:

- `ssh-shell-proxy-x64.exe` — for Windows on x86_64
- `ssh-shell-proxy-arm64.exe` — for Windows on ARM64

### Manual build

```bash
GOOS=windows GOARCH=amd64 go build -o ssh-shell-proxy-x64.exe .
GOOS=windows GOARCH=arm64 go build -o ssh-shell-proxy-arm64.exe .
```
