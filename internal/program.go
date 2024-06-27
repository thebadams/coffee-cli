package program

import tea "github.com/charmbracelet/bubbletea"

type program struct {
	programState int
	programPages []string
}

func DefaultProgram() program {
	return program{
		programPages: []string{"main menu", "new coffee"},
	}
}

func (p program) Init() tea.Cmd {
	return nil
}

func (p program) View() string {
	var s string
	switch p.programState {
	case 0:
		s += p.programPages[0] + "\n\n" + "(press n to create new coffee)"

	case 1:
		s += p.programPages[1] + "\n\n(press m to return to main menu)"
	}
	s += "\n\n(Press q or esc to quit)"
	return s
}

func (p program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return p, tea.Quit
		case "n":
			p.programState = 1
		case "m":
			p.programState = 0
		}
	}
	return p, nil
}
