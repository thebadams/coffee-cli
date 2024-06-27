package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	coffees []string
	cursor  int
}

func initialModel() model {
	return model{
		coffees: []string{"Landrace", "Guji"},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.coffees)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Coffees\n\n"
	for i, coffee := range m.coffees {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += cursor + coffee + "\n"
	}

	s += "\nPress q to quit program"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error Starting Program: %v", err)
		os.Exit(1)
	}
}
