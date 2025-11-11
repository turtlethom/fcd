package internal

import (
	"strings"
	"fmt"
	"os"
	"path/filepath"
)

// IsInsideHome checks if a given path resides within the user's home directory
func IsInsideHome(path string) (string, error) {
	// Expand ~ if present
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(home, path[1:])
	}

	// Convert to absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	// Resolve symlinks
	absPath, err = filepath.EvalSymlinks(absPath)
	if err != nil {
		return "", err
	}

	// Grab the user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	rel, err := filepath.Rel(home, absPath)
	if err != nil || strings.HasPrefix(rel, "..") {
		return "", fmt.Errorf("path %q is outside your home directory", absPath)
	}

	return absPath, nil
}

// SelectColor is a helper function for color selection by the user
func SelectColor(color string) (string, error) {
	switch strings.ToLower(color) {
	case "crimson":
		return "#DC143C", nil
	case "coral":
		return "#F08080", nil
	case "pink": 
		return "#FF007F", nil
	case "red": 
		return "#E06C75", nil
	case "orange":
		return "#D19A66", nil
	case "tangerine":
		return "#FF9500", nil
	case "amber":
		return "#FFBF00", nil
	case "gold":
		return "#DAA520", nil
	case "yellow":
		return "#E5C078", nil
	case "green":
		return "#98C379", nil
	case "emerald":
		return "#50C878", nil
	case "aqua":
		return "#56B6C2", nil
	case "blue":
		return "#1E90FF", nil
	case "navy":
		return "#000080", nil
	case "amethyst":
		return "#9966CC", nil
	case "indigo":
		return "#4B0082", nil
	case "violet":
		return "#C678DD", nil
	case "gray":
		return "#5C6370", nil
	case "white":
		return "#FFFFFF", nil
	case "black":
		return "#000000", nil
	case "brown":
		return "#A0522D", nil
	case "none":
		return "", nil
	default:
		return "", fmt.Errorf("fcd: Invalid color argument %q\n", color)
	}
}
