package opener

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// BuildURL converts a file path or URL argument into a proxy URL.
// For local files, relative paths are resolved to absolute paths.
// For remote URLs (http:// or https://), the scheme is converted to a path prefix.
func BuildURL(arg string, port int) (string, error) {
	base := fmt.Sprintf("http://localhost:%d", port)

	if strings.HasPrefix(arg, "http://") {
		return base + "/http/" + strings.TrimPrefix(arg, "http://"), nil
	}
	if strings.HasPrefix(arg, "https://") {
		return base + "/https/" + strings.TrimPrefix(arg, "https://"), nil
	}

	// Local file path
	absPath, err := filepath.Abs(arg)
	if err != nil {
		return "", fmt.Errorf("failed to resolve path %q: %w", arg, err)
	}
	// Convert to forward slashes for URL (Windows compatibility)
	urlPath := filepath.ToSlash(absPath)
	return base + "/local/" + strings.TrimPrefix(urlPath, "/"), nil
}

// IsServerRunning checks if a server is already listening on the given port.
func IsServerRunning(port int) bool {
	addr := fmt.Sprintf("localhost:%d", port)
	conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// StartServer launches a new server process in the background and waits
// for it to become ready. The provided args should be the command-line
// flags to pass to the server (excluding the file argument).
func StartServer(port int, args []string) error {
	exe, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	cmd := exec.Command(exe, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Detach the child process so it survives after this process exits.
	setSysProcAttr(cmd)

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	// Wait for the server to become ready (up to 5 seconds).
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if IsServerRunning(port) {
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("server did not start within 5 seconds")
}

// OpenBrowser opens the given URL in the default browser.
func OpenBrowser(rawURL string) error {
	// Validate that it's a proper URL before passing to the shell.
	if _, err := url.Parse(rawURL); err != nil {
		return fmt.Errorf("invalid URL %q: %w", rawURL, err)
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", rawURL)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", rawURL)
	case "darwin":
		cmd = exec.Command("open", rawURL)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	return cmd.Start()
}
