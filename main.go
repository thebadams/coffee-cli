package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	coffees []string
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
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Coffees\n\n"
	for _, coffee := range m.coffees {
		s += coffee + "\n"
	}

	s += "\nPress q to quite program"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error Starting Program: %v", err)
		os.Exit(1)
	}
}
