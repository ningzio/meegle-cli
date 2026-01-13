package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// InitConfig initializes the configuration using Viper.
// It looks for config.yaml in ~/.meegle/ and supports environment variables.
func InitConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(home, ".meegle")
	// Ensure the config directory exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.MkdirAll(configPath, 0755); err != nil {
			return err
		}
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.SetEnvPrefix("meegle")
	viper.AutomaticEnv()

	// Defaults can be set here if needed

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired, or create default
			// For now, we just proceed as env vars might be used
		} else {
			// Config file was found but another error produced
			return fmt.Errorf("fatal error config file: %w", err)
		}
	}

	return nil
}

// GetString returns a string value from the config.
func GetString(key string) string {
	return viper.GetString(key)
}
