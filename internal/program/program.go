package program

import tea "github.com/charmbracelet/bubbletea"

type program struct {
	programState int
	programPages []string
}

func defaultProgram() program {
	return program{
		programPages: []string{"main"},
	}
}

func (p program) Init() tea.Cmd {
	return nil
}

func (p program) View() string {
	var s string
	switch p.programState {
	case 0:
		s += "MAIN MENU/n/n"
	}
	return s
}

func (p program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return p, tea.Quit
		}
	}
	return p, nil
}
