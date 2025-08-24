package cmd

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) RenderTitle() string {
	title := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).
		Render("FCD - Shortcut Menu")
	return title + "\n\n"
}

func (m Model) RenderChoices() string {
	var out string
	for i, choice := range m.Choices {
		// Symbol for hovered/unhovered
		symbol := "  ) "
		if m.Cursor == i {
			symbol = " *) "
		}

		// Style for the label (highlighted if hovered)
		labelStyle := lipgloss.NewStyle().
			// Padding(0, 1).
			Bold(true)
		if m.Cursor == i {
			labelStyle = labelStyle.
				// Background(lipgloss.Color("62")).  // teal label background
				Background(lipgloss.Color("38")).  // blue label background
				Foreground(lipgloss.Color("230")) // bright label text
		} else {
			labelStyle = labelStyle.Foreground(lipgloss.Color("250")) // dim label text
		}

		// Render label with brackets
		renderedLabel := labelStyle.Render(fmt.Sprintf("[ %s ]", choice.Label))

		// Style for path (always dim)
		pathStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

		// Combine symbol + label + optional path
		line := fmt.Sprintf("%s%s", symbol, renderedLabel)
		if m.Cursor == i {
			line += fmt.Sprintf(" → %s", pathStyle.Render(choice.Path))
		}

		out += line + "\n"
	}
	return out
}

func (m Model) RenderHelp() string {
	return "(j/k) to move, Enter to select, q to quit.\n"
}
