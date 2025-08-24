package cmd

import (
	"os"

	"github.com/charmbracelet/lipgloss"
)

// Styles bundles all Lipgloss styles
type Styles struct {
	Title       lipgloss.Style
	NormalLabel lipgloss.Style
	ActiveLabel lipgloss.Style
	Path        lipgloss.Style
}

// NewStylesForStderr returns styles bound to stderr
func NewStylesForStderr() Styles {
	r := lipgloss.NewRenderer(os.Stderr)

	return Styles{
		Title:       r.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1),
		NormalLabel: r.NewStyle().Foreground(lipgloss.Color("250")).Bold(true),
		ActiveLabel: r.NewStyle().Background(lipgloss.Color("38")).Foreground(lipgloss.Color("230")).Bold(true),
		Path:        r.NewStyle().Foreground(lipgloss.Color("240")),
	}
}
