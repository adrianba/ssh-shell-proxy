# ssh-shell-proxy

A Windows executable that acts as a default shell for OpenSSH Server, proxying
connections into a WSL Debian instance. When you SSH into the Windows machine,
you land directly in a Linux shell instead of `cmd.exe` or PowerShell.

## How It Works

- **Interactive SSH session** (`ssh user@host`) — opens an interactive WSL Debian
  shell in your Linux home directory.
- **Remote command** (`ssh user@host ps -aux`) — runs the command inside WSL Debian
  via `sh -c`, handling both Linux and Windows SSH client argument styles.
- **Unsupported arguments** — prints an error to stderr.

## Windows Setup

### 1. Install OpenSSH Server

OpenSSH Server is available as a Windows optional feature.

**Via Settings:**

Settings → System → Optional Features → Add a feature → **OpenSSH Server** → Install

**Via PowerShell (as Administrator):**

```powershell
Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0
```

Then start and enable the service:

```powershell
Start-Service sshd
Set-Service -Name sshd -StartupType Automatic
```

### 2. Install ssh-shell-proxy

Copy the appropriate `.exe` to a permanent location on the Windows machine, for example:

```powershell
Copy-Item ssh-shell-proxy-x64.exe C:\Program` Files\ssh-shell-proxy\ssh-shell-proxy.exe
```

Use `ssh-shell-proxy-arm64.exe` for ARM64 devices (e.g. Surface Pro X).

### 3. Configure as Default Shell

Set ssh-shell-proxy as the default shell for OpenSSH Server (PowerShell as Administrator):

```powershell
New-ItemProperty -Path "HKLM:\SOFTWARE\OpenSSH" `
    -Name DefaultShell `
    -Value "C:\Program Files\ssh-shell-proxy\ssh-shell-proxy.exe" `
    -PropertyType String -Force
```

Restart the SSH service:

```powershell
Restart-Service sshd
```

Now when you SSH into the machine, you'll get a WSL Debian shell automatically.

## Building from Source

### Prerequisites — Install Go

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

### Build

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

## Releasing

Releases are published automatically via GitHub Actions when a version tag is pushed.

To create a new release:

```bash
git tag v0.2
git push origin v0.2
```

This triggers the release workflow which builds both Windows binaries and publishes
them as a GitHub release. The version number from the tag is embedded into the
binaries and can be checked with `ssh-shell-proxy.exe --version`.

You can also pass a version to the build script for local builds:

```bash
./build.sh 0.2
```
