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

// Shortcut stores the corresponding Label and Path to a bookmarked directory,
// as defined in fcd_config.json
// It is also used as an Item component for the List (bubbles/list)
type Shortcut struct {
	Label string `json:"label"`		// Represents the label provided by user
	Path  string `json:"path"`		// Represents the path provided by user
}

// Title returns the Item's title for display in bubbles/list component
// Implements list.Item interface
func (s Shortcut) Title() string       { return s.Label }

// Description returns the Item's description for display in bubbles/list component
// Implements list.Item interface
func (s Shortcut) Description() string { return s.Path }

// FilterValue returns the Item's filtered value in bubbles/list component
// Implements list.Item interface
func (s Shortcut) FilterValue() string { return s.Label }

// UserColors defines the user's colors for the bubbles/list component,
// as defined within fcd_config.json
type UserColors struct {
	Primary		string `json:"primary"`
	Secondary	string `json:"secondary"`
	Tertiary  string `json:"tertiary"`
}

// Config defines the user's configuration settings for fcd,
// as defined with fcd_config.json
type Config struct {
	Shortcuts 	[]Shortcut 	`json:"shortcuts"`
	UserColors 	UserColors 	`json:"userColors"`
}

// fetchUserConfig returns the absolute path of fcd_config.json,
// based on the user's operating system
func fetchUserConfig() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "fcd", "fcd_config.json"), nil	// [USERCONFIG]/fcd/fcd_config.json
}

// CreateConfig initializes the configuration for fcd_config.json,
// and returns a pointer to a new Config
// 
// When invoked,
func CreateConfig(path string) (*Config, error) {
	config := &Config{
		Shortcuts: []Shortcut{},
		UserColors: UserColors{
				Primary: "",
				Secondary: "",
				Tertiary: "",
		},
	}
	if err := SaveToConfig(config); err != nil {
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
	// If empty file or bad JSON → reset to empty config
	if err := decoder.Decode(&config); err != nil {
		return &Config{Shortcuts: []Shortcut{}}, nil
	}
	return &config, nil
}

// Loads config or creates it if missing
func HandleConfig() (*Config, error) {
	path, err := fetchUserConfig()
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

// SaveToConfig stores the new state of the config back to json file
func SaveToConfig(config *Config) error {
	path, err := fetchUserConfig()
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
