package internal

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Shortcut struct {
	Label string `json:"label"`
	Path  string `json:"path"`
}

type Config struct {
	Shortcuts []Shortcut `json:"shortcuts"`
}

// getConfigPath returns absolute path to fcd_config.json
func getConfigPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "fcd", "fcd_config.json"), nil
}

// CreateConfig makes a brand new config file if it doesn't exist
func CreateConfig(path string) (*Config, error) {
	config := &Config{Shortcuts: []Shortcut{}}

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}

	// Save empty config
	if err := save(config, path); err != nil {
		return nil, err
	}
	return config, nil
}

// Reads the JSON config from disk
func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var config Config
	if err := json.NewDecoder(f).Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

// Loads config or creates it if missing
func HandleConfig() (*Config, error) {
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return CreateConfig(path)
	}
	return LoadConfig(path)
}

// Save writes the config back to disk
func Save(config *Config) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}
	return save(config, path)
}

// internal helper for saving with a specific path
func save(config *Config, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(config)
}
