// ssh-shell-proxy launches a WSL Debian shell from Windows.
// It acts like a shell that supports "-c" to run a single command,
// or opens an interactive session when called with no arguments.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// wslPath is the full path to the Windows WSL executable.
// Using the absolute path avoids relying on PATH resolution.
const wslPath = `C:\Windows\System32\wsl.exe`

// version is the program version, set at build time via -ldflags.
var version = "dev"

func main() {
	// run() returns an exit code; we pass it to os.Exit so that
	// callers (e.g. SSH, scripts) see the correct status.
	os.Exit(run())
}

func run() int {
	// os.Args[0] is the program name itself, so [1:] gives us
	// only the arguments the user actually passed.
	args := os.Args[1:]

	// No arguments: open an interactive WSL Debian shell
	// starting in the user's Linux home directory (~).
	if len(args) == 0 {
		return execWSL("--distribution", "Debian", "--cd", "~")
	}

	// Show version and exit.
	if args[0] == "-v" || args[0] == "-V" || args[0] == "--version" {
		fmt.Printf("ssh-shell-proxy %s\n", version)
		return 0
	}

	// "-c <command>" mode: behaves like "sh -c" — runs a command
	// inside WSL and exits.
	if args[0] == "-c" {
		remainder := args[1:]
		if len(remainder) == 0 {
			fmt.Fprintln(os.Stderr, "error: -c requires a command")
			return 1
		}
		// Join all arguments into a single command string and pass
		// it to "sh -c" inside WSL. This ensures correct behavior
		// regardless of whether the caller sends the command as
		// separate args (Linux SSH) or a single string (Windows SSH).
		command := strings.Join(remainder, " ")
		return execWSL("--distribution", "Debian", "--cd", "~", "--", "sh", "-c", command)
	}

	// Anything else is unsupported — report what we received so
	// the user can see what went wrong.
	fmt.Fprintf(os.Stderr, "error: unsupported arguments: %s\n", strings.Join(args, " "))
	return 1
}

// execWSL runs wsl.exe with the given arguments, wiring up stdin/stdout/stderr
// so the WSL process can interact with the terminal directly.
// It returns the exit code from wsl.exe (or 1 if something else goes wrong).
func execWSL(args ...string) int {
	cmd := exec.Command(wslPath, args...)

	// Connect the child process to our own terminal streams so that
	// interactive programs (vim, bash, etc.) work correctly.
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command and wait for it to finish.
	if err := cmd.Run(); err != nil {
		// If wsl.exe exited with a non-zero code, propagate that
		// exact code back to our caller.
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode()
		}
		// Some other error (e.g. wsl.exe not found) — report it.
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return 1
	}
	return 0
}
