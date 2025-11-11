package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
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
// as defined within fcd_config.json
//
// It manages the state of the current state within the configuration
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
// fcdConfigPath - String representing the absolute path to fcd_config.json
func CreateDefaultConfig() (*Config, error) {
	config := &Config{
		Shortcuts: []Shortcut{},
		UserColors: UserColors{
				Primary: "",
				Secondary: "",
				Tertiary: "",
		},
	}
	// Saving default state of new Config to fcd_config.json
	if err := SaveToConfig(config); err != nil {
		return nil, err
	}
	return config, nil
}

// LoadConfig reads data from fcd_config.json, decoding user configuration
// settings into a new Config
//
// fcdConfigPath - String representing the absolute path to fcd_config.json
func LoadConfig(fcdConfigPath string) (*Config, error) {
	configFile, err := os.Open(fcdConfigPath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config Config
	decoder := json.NewDecoder(configFile)
	// If file is empty or malformed JSON data in config file → reset to empty config
	if err := decoder.Decode(&config); err != nil {
		// return &Config{Shortcuts: []Shortcut{}}, nil
		return CreateDefaultConfig();
	}
	return &config, nil
}

// HandleConfig controls whether a user's fcd_config.json must be initialized or loaded
//
// fcd_config.json is initialized with default configuration settings if not present
// Else, fcd_config.json is loaded into Config instance
func HandleConfig() (*Config, error) {
	fcdConfigPath, err := fetchUserConfig()
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(fcdConfigPath); errors.Is(err, os.ErrNotExist) {
		return CreateDefaultConfig()
	}
	return LoadConfig(fcdConfigPath)
}
// AddShortcut appends a new shortcut to the Config.Shortcuts array of a Config instance
//
// label - String representing title for new Shortcut
// dirPath - String representing directory path for new Shortcut
func (c *Config) AddShortcut(label, dirPath string) error {
	// Ensures the path to the shortcut is within the user's home directory
	absPath, err := IsInsideHome(dirPath)
	if err != nil {
		return err
	}
	// Checks whether a duplicate Shortcut exists within Config.Shortcuts
	for _, sc := range c.Shortcuts {
		if sc.Label == label {
			return fmt.Errorf("shortcut with label %q already exists", label)
		}
		if sc.Path == absPath {
			return fmt.Errorf("shortcut for path %q already exists", absPath)
		}
	}
	// Appends new Shortcut to Config.Shortcuts
	c.Shortcuts = append(c.Shortcuts, Shortcut{
		Label: label,
		Path:  absPath,
	})
	return nil
}

// SaveToConfig stores the current state of the Config back to fcd_config.json file
func SaveToConfig(config *Config) error {
	fcdConfigPath, err := fetchUserConfig()
	if err != nil {
		return err
	}
	// Formatting JSON data from fcd_config.json
	fcdConfigData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	// Ensure parent directory exists
	if err := os.MkdirAll(filepath.Dir(fcdConfigPath), 0o755); err != nil {
		return err
	}
	// Write Config state to fcd_config.json
	return os.WriteFile(fcdConfigPath, fcdConfigData, 0o644)
}

// ToListModel parses Config.Shortcuts from user configuration into list.Model
func (c *Config) ToListModel() list.Model {
	items := make([]list.Item, 0, len(c.Shortcuts))
	for i, sc := range c.Shortcuts {
		items[i] = sc
	}
	listModel := list.New(items, list.NewDefaultDelegate(), 0, 0)
	listModel.Title = "Shortcuts"
	return listModel
}
