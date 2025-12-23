//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	binaryName = "api"
	mainPath   = "./cmd/api"
	cliName    = "gostyl"
	cliPath    = "./cmd/gostyl"
)

// Values injected at build time
var (
	Version = "dev"
	Commit  = "none"
)

var Default = Build

// Build creates an optimized release binary for the API
func Build() error {
	return build(binaryName, mainPath)
}

// BuildCLI creates an optimized release binary for the CLI
func BuildCLI() error {
	return build(cliName, cliPath)
}

func build(out string, path string) error {
	fmt.Printf("🔨 Building %s...\n", out)

	ldflags := fmt.Sprintf(
		"-s -w -X main.version=%s -X main.commit=%s",
		Version,
		Commit,
	)

	cmd := exec.Command(
		"go", "build",
		"-trimpath",
		"-ldflags", ldflags,
		"-o", out,
		path,
	)

	cmd.Env = append(os.Environ(),
		"CGO_ENABLED=0",
		fmt.Sprintf("GOOS=%s", runtime.GOOS),
		fmt.Sprintf("GOARCH=%s", runtime.GOARCH),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// Run executes the built binary (production-like)
func Run() error {
	fmt.Println("🚀 Running API...")
	if err := Build(); err != nil {
		return err
	}

	cmd := exec.Command("./" + binaryName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	return cmd.Run()
}

// Dev runs the API without building (fast dev loop)
func Dev() error {
	fmt.Println("🧪 Running in dev mode...")
	cmd := exec.Command("go", "run", mainPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	return cmd.Run()
}

// Test runs all tests
func Test() error {
	fmt.Println("🧪 Running tests...")
	cmd := exec.Command("go", "test", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Clean removes the built binary
func Clean() error {
	fmt.Println("🧹 Cleaning...")
	return os.RemoveAll(binaryName)
}
