package program

import (
	tea "github.com/charmbracelet/bubbletea"
)

func CreateProgram() tea.Program {
	return *tea.NewProgram(rootProgram())
}
