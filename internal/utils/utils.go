package utils

import (
	"strings"
	"fmt"
	"os"
	"path/filepath"
)

func IsInsideHome(path string) (string, error) {
	// Expand ~ if present
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(home, path[1:])
	}

	// Convert to absolute
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
