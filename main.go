// Cycloid CLI installer shim.
//
// Install via: go install github.com/cycloidio/cycloid-cli@latest
//
// On first run, this binary downloads the correct platform-specific pre-built
// cy binary from GitHub Releases, caches it locally, and exec's it with all
// arguments passed through. Subsequent runs exec the cached binary directly.
//
// Set CY_VERSION to pin a specific version (e.g., CY_VERSION=v6.10.24).
// The default version is the one this shim was built with.
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

// version is set at build time via ldflags.
var version = "dev"

const (
	releasesBaseURL = "https://github.com/cycloidio/cycloid-cli/releases/download"
	cacheDir        = ".cycloid/bin"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--shim-version" {
		fmt.Printf("cycloid-cli shim %s (%s/%s)\n", version, runtime.GOOS, runtime.GOARCH)
		os.Exit(0)
	}

	binPath, err := ensureBinary()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cycloid-cli: failed to get cy binary: %v\n", err)
		os.Exit(1)
	}

	if err := execBinary(binPath, os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "cycloid-cli: %v\n", err)
		os.Exit(1)
	}
}

func ensureBinary() (string, error) {
	v := resolveVersion()
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot determine home directory: %w", err)
	}

	ext := ""
	if runtime.GOOS == "windows" {
		ext = ".exe"
	}
	binName := fmt.Sprintf("cy-%s-%s%s", runtime.GOOS, runtime.GOARCH, ext)
	cachedPath := filepath.Join(home, cacheDir, fmt.Sprintf("cy-%s%s", v, ext))

	if _, err := os.Stat(cachedPath); err == nil {
		return cachedPath, nil
	}

	url := fmt.Sprintf("%s/%s/%s", releasesBaseURL, v, binName)
	fmt.Fprintf(os.Stderr, "cycloid-cli: downloading %s ...\n", url)

	if err := downloadFile(cachedPath, url); err != nil {
		return "", err
	}

	if err := os.Chmod(cachedPath, 0o755); err != nil {
		return "", fmt.Errorf("chmod: %w", err)
	}

	checksumURL := fmt.Sprintf("%s/%s/checksums.txt", releasesBaseURL, v)
	if err := verifyChecksum(cachedPath, binName, checksumURL); err != nil {
		fmt.Fprintf(os.Stderr, "cycloid-cli: warning: checksum verification skipped: %v\n", err)
	}

	return cachedPath, nil
}

func resolveVersion() string {
	if v := os.Getenv("CY_VERSION"); v != "" {
		return v
	}
	return version
}

func downloadFile(dst, url string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	resp, err := http.Get(url) //nolint:gosec // URL is constructed from constants
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: HTTP %d from %s", resp.StatusCode, url)
	}

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		os.Remove(dst)
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

func verifyChecksum(filePath, binName, checksumURL string) error {
	resp, err := http.Get(checksumURL) //nolint:gosec
	if err != nil || resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not fetch checksums")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read checksums: %w", err)
	}

	var expectedHash string
	for _, line := range strings.Split(string(body), "\n") {
		parts := strings.Fields(line)
		if len(parts) == 2 && parts[1] == binName {
			expectedHash = parts[0]
			break
		}
	}
	if expectedHash == "" {
		return fmt.Errorf("no checksum found for %s", binName)
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return err
	}

	actualHash := hex.EncodeToString(h.Sum(nil))
	if actualHash != expectedHash {
		os.Remove(filePath)
		return fmt.Errorf("checksum mismatch: expected %s, got %s", expectedHash, actualHash)
	}

	return nil
}

func execBinary(binPath string, args []string) error {
	if runtime.GOOS == "windows" {
		cmd := exec.Command(binPath, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	argv := append([]string{binPath}, args...)
	return syscall.Exec(binPath, argv, os.Environ())
}
