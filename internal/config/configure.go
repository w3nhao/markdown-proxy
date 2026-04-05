package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// RunConfigure runs the interactive configuration wizard.
// It reads from r and writes prompts to w, then saves the result to the config file.
func RunConfigure(r io.Reader, w io.Writer) error {
	// Load existing config as defaults
	fc := loadConfigFile()

	scanner := bufio.NewScanner(r)

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "markdown-proxy configuration")
	fmt.Fprintln(w, "=============================")
	fmt.Fprintln(w, "")

	// PlantUML server
	fmt.Fprintln(w, "PlantUML server URL (empty to disable):")
	fmt.Fprintln(w, "  WARNING: Diagram content will be sent to the specified server.")
	fmt.Fprintln(w, "  Use a self-hosted server for sensitive diagrams.")
	fmt.Fprintln(w, "  Public server: https://www.plantuml.com/plantuml")
	fc.PlantUMLServer = promptString(scanner, w, fc.PlantUMLServer)
	fmt.Fprintln(w, "")

	// Theme
	defaultTheme := fc.Theme
	if defaultTheme == "" {
		defaultTheme = "github"
	}
	fmt.Fprintf(w, "Default theme (github/simple/dark) [%s]:\n", defaultTheme)
	val := promptString(scanner, w, defaultTheme)
	if val == "" {
		val = defaultTheme
	}
	fc.Theme = val
	fmt.Fprintln(w, "")

	// Port
	defaultPort := fc.Port
	if defaultPort == 0 {
		defaultPort = 9080
	}
	fmt.Fprintf(w, "Listen port [%d]:\n", defaultPort)
	portStr := promptString(scanner, w, strconv.Itoa(defaultPort))
	if p, err := strconv.Atoi(portStr); err == nil && p > 0 {
		fc.Port = p
	} else {
		fc.Port = defaultPort
	}
	fmt.Fprintln(w, "")

	// Listen address
	defaultListen := fc.Listen
	if defaultListen == "" {
		defaultListen = "127.0.0.1"
	}
	fmt.Fprintf(w, "Listen address [%s]:\n", defaultListen)
	val = promptString(scanner, w, defaultListen)
	if val == "" {
		val = defaultListen
	}
	fc.Listen = val
	fmt.Fprintln(w, "")

	// Save
	path, err := configFilePathFunc()
	if err != nil {
		return fmt.Errorf("cannot determine config path: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("cannot create config directory: %w", err)
	}

	data, err := json.MarshalIndent(fc, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal config: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("cannot write config file: %w", err)
	}

	fmt.Fprintf(w, "Configuration saved to %s\n", path)
	return nil
}

// promptString displays "> " and reads one line. If the input is empty, returns defaultVal.
func promptString(scanner *bufio.Scanner, w io.Writer, defaultVal string) string {
	if defaultVal != "" {
		fmt.Fprintf(w, "> [%s] ", defaultVal)
	} else {
		fmt.Fprint(w, "> ")
	}
	if scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			return line
		}
	}
	return defaultVal
}
