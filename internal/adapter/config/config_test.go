package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"meegle-cli/internal/adapter/config"
)

func TestInitConfig_EnvVars(t *testing.T) {
	// Setup env var
	t.Setenv("MEEGLE_TEST_KEY", "env_value")

	// Reset viper
	viper.Reset()

	// Init config
	if err := config.InitConfig(); err != nil {
		t.Fatalf("InitConfig failed: %v", err)
	}

	// Verify env var reading
	if val := config.GetString("test_key"); val != "env_value" {
		t.Errorf("Expected 'env_value', got '%s'", val)
	}
}

func TestInitConfig_File(t *testing.T) {
	// Setup temp home dir
	tempHome := t.TempDir()
	t.Setenv("HOME", tempHome)

	// Create config file
	configDir := filepath.Join(tempHome, ".meegle")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Fatal(err)
	}
	configFile := filepath.Join(configDir, "config.yaml")
	content := []byte("file_key: file_value")
	if err := os.WriteFile(configFile, content, 0644); err != nil {
		t.Fatal(err)
	}

	// Reset viper
	viper.Reset()

	// Init config
	if err := config.InitConfig(); err != nil {
		t.Fatalf("InitConfig failed: %v", err)
	}

	// Verify file reading
	if val := config.GetString("file_key"); val != "file_value" {
		t.Errorf("Expected 'file_value', got '%s'", val)
	}
}
