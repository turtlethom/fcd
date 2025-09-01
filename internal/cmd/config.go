package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	"github.com/turtlethom/fcd/internal/utils"
)

/*
	A Shortcut holds the data related to an entry by the user, and is prepared
	to be written to the user's Config:
		- `label` is the label for the filepath
		- `path` is the filepath within the user's home directory
*/
type Shortcut struct {
	Label string `json:"label"`
	Path  string `json:"path"`
}

func (s Shortcut) Title() string       { return s.Label }
func (s Shortcut) Description() string { return s.Path }
func (s Shortcut) FilterValue() string { return s.Label }

/*
	Config holds all the shortcuts for a given user
		- `Shortcuts` holds all data related to the user's valid shortcuts
*/
type Config struct {
	Shortcuts []Shortcut `json:"shortcuts"`
}

// Returns absolute path to fcd_config.json
func getConfigPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "fcd", "fcd_config.json"), nil
}

// Makes a brand new config file if it doesn't exist
func CreateConfig(path string) (*Config, error) {
	config := &Config{Shortcuts: []Shortcut{}}
	if err := Save(config); err != nil {
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
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&config); err != nil {
		// If empty file or bad JSON → reset to empty config
		return &Config{Shortcuts: []Shortcut{}}, nil
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
// Adds a new shortcut if its label AND path are unique.
func (c *Config) AddShortcut(label, path string) error {
	absPath, err := utils.IsInsideHome(path)
	if err != nil {
		return err
	}

	// Check for duplicates
	for _, sc := range c.Shortcuts {
		if sc.Label == label {
			return fmt.Errorf("shortcut with label %q already exists", label)
		}
		if sc.Path == absPath {
			return fmt.Errorf("shortcut for path %q already exists", absPath)
		}
	}

	c.Shortcuts = append(c.Shortcuts, Shortcut{
		Label: label,
		Path:  absPath,
	})
	return nil
}

// Save new state of the config back to json file
func Save(config *Config) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	// Formatting JSON data
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	// Check if parent directory exists
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}

// Parse shortcuts from user config into list.Model
func (c *Config) ToListModel() list.Model {
	items := make([]list.Item, 0, len(c.Shortcuts))
	for i, sc := range c.Shortcuts {
		items[i] = sc
	}
	m := list.New(items, list.NewDefaultDelegate(), 0, 0)
	m.Title = "Shortcuts"
	return m
}
