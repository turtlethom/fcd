package cmd

import (
	"fmt"
	"strings"
)

func (m Model) RenderTitle() string {
	return m.Styles.Title.Render("FCD - [ Choose a shortcut ]:") + "\n"
}

func (m Model) RenderChoices() string {
	var output strings.Builder
	for i, choice := range m.Choices {
		// NORMAL CURSOR SYMBOL
		symbol := "  ) "
		style := m.Styles.NormalLabel
		if m.Cursor == i {
			// ACTIVE CURSOR SYMBOL
			symbol = " *) "
			style = m.Styles.ActiveLabel
		}
		// RENDER LABEL
		label := style.Render(fmt.Sprintf("[ %s ]", choice.Label))
		line := fmt.Sprintf("%s%s", symbol, label)
		if m.Cursor == i {
			line += fmt.Sprintf(" → %s", m.Styles.Path.Render(choice.Path))
		}
		fmt.Fprintln(&output, line)
	}
	return output.String()
}

func (m Model) RenderHelp() string {
	return "(j/k) to move, Enter to select, q to quit.\n"
}
