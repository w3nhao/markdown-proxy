package config

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunConfigure(t *testing.T) {
	// Use a temp dir as config dir
	tmpDir := t.TempDir()
	origFunc := configFilePathFunc
	configFilePathFunc = func() (string, error) {
		return filepath.Join(tmpDir, "config.json"), nil
	}
	defer func() { configFilePathFunc = origFunc }()

	input := "https://www.plantuml.com/plantuml\ndark\n8080\n0.0.0.0\n"
	var out bytes.Buffer

	err := RunConfigure(strings.NewReader(input), &out)
	if err != nil {
		t.Fatalf("RunConfigure() error = %v", err)
	}

	// Verify output contains expected prompts
	output := out.String()
	if !strings.Contains(output, "markdown-proxy configuration") {
		t.Error("output should contain header")
	}
	if !strings.Contains(output, "WARNING: Diagram content will be sent") {
		t.Error("output should contain PlantUML warning")
	}
	if !strings.Contains(output, "Configuration saved to") {
		t.Error("output should contain save confirmation")
	}

	// Verify saved file
	data, err := os.ReadFile(filepath.Join(tmpDir, "config.json"))
	if err != nil {
		t.Fatalf("cannot read config file: %v", err)
	}

	var fc fileConfig
	if err := json.Unmarshal(data, &fc); err != nil {
		t.Fatalf("cannot parse config file: %v", err)
	}

	if fc.PlantUMLServer != "https://www.plantuml.com/plantuml" {
		t.Errorf("PlantUMLServer = %q, want %q", fc.PlantUMLServer, "https://www.plantuml.com/plantuml")
	}
	if fc.Theme != "dark" {
		t.Errorf("Theme = %q, want %q", fc.Theme, "dark")
	}
	if fc.Port != 8080 {
		t.Errorf("Port = %d, want %d", fc.Port, 8080)
	}
	if fc.Listen != "0.0.0.0" {
		t.Errorf("Listen = %q, want %q", fc.Listen, "0.0.0.0")
	}
}

func TestRunConfigure_Defaults(t *testing.T) {
	// Use a temp dir as config dir
	tmpDir := t.TempDir()
	origFunc := configFilePathFunc
	configFilePathFunc = func() (string, error) {
		return filepath.Join(tmpDir, "config.json"), nil
	}
	defer func() { configFilePathFunc = origFunc }()

	// All empty inputs → use defaults
	input := "\n\n\n\n"
	var out bytes.Buffer

	err := RunConfigure(strings.NewReader(input), &out)
	if err != nil {
		t.Fatalf("RunConfigure() error = %v", err)
	}

	data, err := os.ReadFile(filepath.Join(tmpDir, "config.json"))
	if err != nil {
		t.Fatalf("cannot read config file: %v", err)
	}

	var fc fileConfig
	if err := json.Unmarshal(data, &fc); err != nil {
		t.Fatalf("cannot parse config file: %v", err)
	}

	if fc.PlantUMLServer != "" {
		t.Errorf("PlantUMLServer = %q, want empty", fc.PlantUMLServer)
	}
	if fc.Theme != "github" {
		t.Errorf("Theme = %q, want %q", fc.Theme, "github")
	}
	if fc.Port != 9080 {
		t.Errorf("Port = %d, want %d", fc.Port, 9080)
	}
	if fc.Listen != "127.0.0.1" {
		t.Errorf("Listen = %q, want %q", fc.Listen, "127.0.0.1")
	}
}

func TestRunConfigure_ExistingConfig(t *testing.T) {
	tmpDir := t.TempDir()
	origFunc := configFilePathFunc
	configFilePathFunc = func() (string, error) {
		return filepath.Join(tmpDir, "config.json"), nil
	}
	defer func() { configFilePathFunc = origFunc }()

	// Write an existing config
	existing := fileConfig{
		PlantUMLServer: "https://plantuml.example.com",
		Theme:          "simple",
		Port:           3000,
		Listen:         "192.168.1.10",
	}
	data, _ := json.Marshal(existing)
	os.WriteFile(filepath.Join(tmpDir, "config.json"), data, 0644)

	// All empty inputs → keep existing values
	input := "\n\n\n\n"
	var out bytes.Buffer

	err := RunConfigure(strings.NewReader(input), &out)
	if err != nil {
		t.Fatalf("RunConfigure() error = %v", err)
	}

	output := out.String()
	// Verify existing values shown as defaults
	if !strings.Contains(output, "https://plantuml.example.com") {
		t.Error("output should show existing PlantUML server as default")
	}
	if !strings.Contains(output, "simple") {
		t.Error("output should show existing theme as default")
	}
	if !strings.Contains(output, "3000") {
		t.Error("output should show existing port as default")
	}

	// Verify values preserved
	data, _ = os.ReadFile(filepath.Join(tmpDir, "config.json"))
	var fc fileConfig
	json.Unmarshal(data, &fc)

	if fc.PlantUMLServer != "https://plantuml.example.com" {
		t.Errorf("PlantUMLServer = %q, want preserved value", fc.PlantUMLServer)
	}
	if fc.Port != 3000 {
		t.Errorf("Port = %d, want 3000", fc.Port)
	}
}

func TestLoadConfigFile(t *testing.T) {
	tmpDir := t.TempDir()
	origFunc := configFilePathFunc
	configFilePathFunc = func() (string, error) {
		return filepath.Join(tmpDir, "config.json"), nil
	}
	defer func() { configFilePathFunc = origFunc }()

	// No file → zero values
	fc := loadConfigFile()
	if fc.PlantUMLServer != "" || fc.Theme != "" || fc.Port != 0 {
		t.Errorf("loadConfigFile() with no file should return zero values, got %+v", fc)
	}

	// Write a config file
	config := `{"plantuml-server": "https://example.com/plantuml", "theme": "dark", "port": 7070}`
	os.WriteFile(filepath.Join(tmpDir, "config.json"), []byte(config), 0644)

	fc = loadConfigFile()
	if fc.PlantUMLServer != "https://example.com/plantuml" {
		t.Errorf("PlantUMLServer = %q, want %q", fc.PlantUMLServer, "https://example.com/plantuml")
	}
	if fc.Theme != "dark" {
		t.Errorf("Theme = %q, want %q", fc.Theme, "dark")
	}
	if fc.Port != 7070 {
		t.Errorf("Port = %d, want %d", fc.Port, 7070)
	}
}

func TestLoadConfigFile_InvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	origFunc := configFilePathFunc
	configFilePathFunc = func() (string, error) {
		return filepath.Join(tmpDir, "config.json"), nil
	}
	defer func() { configFilePathFunc = origFunc }()

	os.WriteFile(filepath.Join(tmpDir, "config.json"), []byte("not json"), 0644)

	fc := loadConfigFile()
	if fc.PlantUMLServer != "" || fc.Theme != "" || fc.Port != 0 {
		t.Errorf("loadConfigFile() with invalid JSON should return zero values, got %+v", fc)
	}
}
