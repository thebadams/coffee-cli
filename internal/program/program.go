package program

import tea "github.com/charmbracelet/bubbletea"

type program struct {
	programState int
	programPages []string
	coffees      []coffee
}

func defaultProgram() program {
	return program{
		programPages: []string{"main menu", "new coffee", "list coffee"},
		coffees: []coffee{
			{"Munyinya Hill, Lot #2", "Sey Coffee", []string{}},
			{"Finca La Estrella - Late Harvest", "Sey Coffee", []string{}},
			{"El Cedro", "Sey Coffee", []string{}},
			{"Sierra Morena", "Sey Coffee", []string{}},
		},
	}
}

func (p program) Init() tea.Cmd {
	return nil
}

func (p program) View() string {
	var s string
	switch p.programState {
	case 0:
		s = mainMenuString(p)
	case 1:
		s = newCoffeeString(p)
	case 2:
		s = listCoffeeString(p)
	}

	s += footerString()
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
		case "l":
			p.programState = 2

		}
	}
	return p, nil
}

func CreateProgram() *tea.Program {
	p := tea.NewProgram(defaultProgram())

	return p
}
